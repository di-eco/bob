package bob

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/benchkram/bob/bobtask"
	"github.com/benchkram/bob/pkg/usererror"

	"github.com/logrusorgru/aurora"

	"github.com/benchkram/errz"

	"github.com/benchkram/bob/bob/bobfile"
	"github.com/benchkram/bob/pkg/filepathutil"
	"github.com/hashicorp/go-version"
)

var (
	ErrDuplicateProjectName = fmt.Errorf("duplicate project name")
)

// find bobfiles recursively.
func (b *B) find() (bobfiles []string, err error) {
	defer errz.Recover(&err)

	list, err := filepathutil.ListRecursive(b.dir)
	errz.Fatal(err)

	for _, file := range list {
		if bobfile.IsBobfile(file) {
			bobfiles = append(bobfiles, file)
		}
	}

	return bobfiles, nil
}

func (b *B) PrintVersionCompatibility(bobfile *bobfile.Bobfile) {
	binVersion, _ := version.NewVersion(Version)

	for _, boblet := range bobfile.Bobfiles() {
		if boblet.Version != "" {
			bobletVersion, _ := version.NewVersion(boblet.Version)

			if binVersion.Core().Segments64()[0] != bobletVersion.Core().Segments64()[0] {
				fmt.Println(aurora.Red(fmt.Sprintf("Warning: major version mismatch: Your bobfile's major version (%s, '%s') is different from the CLI version (%s). This might lead to unexpected errors.", boblet.Version, boblet.Dir(), binVersion)).String())
				continue
			}

			if binVersion.LessThan(bobletVersion) {
				fmt.Println(aurora.Red(fmt.Sprintf("Warning: possible version incompatibility: Your bobfile's version (%s, '%s') is higher than the CLI version (%s). Some features might not work as expected.", boblet.Version, boblet.Dir(), binVersion)).String())
				continue
			}
		}
	}
}

// Aggregate determine and read Bobfiles recursively into memory
// and returns a single Bobfile containing all tasks & runs.
func (b *B) Aggregate() (aggregate *bobfile.Bobfile, err error) {
	defer errz.Recover(&err)

	bobfiles, err := b.find()
	errz.Fatal(err)

	// FIXME: As we don't refer to a child task by projectname but by path
	// it seems to be save to allow duplicate projectnames.
	//projectNames := map[string]bool{}

	// Read & Find Bobfiles
	bobs := []*bobfile.Bobfile{}
	for _, bf := range bobfiles {
		boblet, err := bobfile.BobfileRead(filepath.Dir(bf))
		errz.Fatal(err)

		if boblet.Dir() == b.dir {
			aggregate = boblet
		}

		// FIXME: As we don't refer to a child task by projectname but by path
		// it seems to be save to allow duplicate projectnames.
		//
		// Make sure project names are unique
		// if boblet.Project != "" {
		// 	if ok := projectNames[boblet.Project]; ok {
		// 		return nil, usererror.Wrap(fmt.Errorf("%w found, [%s]", ErrDuplicateProjectName, boblet.Project))
		// 	}
		// 	projectNames[boblet.Project] = true
		// }

		// add env vars and build tasks
		for variable, value := range boblet.Variables {
			for key, task := range boblet.BTasks {
				// TODO: Create and use envvar sanitizer

				task.AddEnvironment(strings.ToUpper(variable), value)

				boblet.BTasks[key] = task
			}
		}

		bobs = append(bobs, boblet)
	}

	if aggregate == nil {
		return nil, usererror.Wrap(ErrCouldNotFindTopLevelBobfile)
	}

	if aggregate.Project == "" {
		// TODO: maybe don't leak absolute path of environment
		aggregate.Project = aggregate.Dir()
	}

	// set project names for all bobfiles and build tasks
	for _, bobfile := range bobs {
		bobfile.Project = aggregate.Project

		for taskname, task := range bobfile.BTasks {
			// Should be the name of the umbrella-bobfile.
			task.SetProject(aggregate.Project)

			// Overwrite value in build map
			bobfile.BTasks[taskname] = task
		}
	}

	aggregate.SetBobfiles(bobs)

	// Merge tasks into one Bobfile
	for _, bobfile := range bobs {
		// Skip the aggregate
		if bobfile.Dir() == aggregate.Dir() {
			continue
		}

		for taskname, task := range bobfile.BTasks {
			dir := bobfile.Dir()

			// Use a relative path as task prefix.
			prefix := strings.TrimPrefix(dir, b.dir)
			taskname := addTaskPrefix(prefix, taskname)

			// Alter the taskname.
			task.SetName(taskname)

			// Rewrite dependent tasks to global scope.
			dependsOn := []string{}
			for _, dependentTask := range task.DependsOn {
				dependsOn = append(dependsOn, addTaskPrefix(prefix, dependentTask))
			}
			task.DependsOn = dependsOn

			aggregate.BTasks[taskname] = task
		}
	}

	// Merge runs into one Bobfile
	for _, bobfile := range bobs {
		// Skip the aggregate
		if bobfile.Dir() == aggregate.Dir() {
			continue
		}

		for runname, run := range bobfile.RTasks {
			dir := bobfile.Dir()

			// Use a relative path as task prefix.
			prefix := strings.TrimPrefix(dir, b.dir)

			runname = addTaskPrefix(prefix, runname)

			// Alter the runname.
			run.SetName(runname)

			// Rewrite dependents to global scope.
			dependsOn := []string{}
			for _, dependent := range run.DependsOn {
				dependsOn = append(dependsOn, addTaskPrefix(prefix, dependent))
			}
			run.DependsOn = dependsOn

			aggregate.RTasks[runname] = run
		}
	}

	// TODO: Gather missing tasks from remote  & Unpack?

	// Gather environment from dependent tasks.
	//
	// Each export is translated into environment variables named:
	//   `second-level/openapi => SECOND_LEVEL_OPENAPI`
	// hyphens`-` are translated to underscores`_`.
	//
	// The file is prefixed with all paths to make it relative to dir of the the top Bobfile:
	//   `openapi.yaml => sencond-level/openapi.yaml`
	//
	// TODO: Exports should be part of a packed file and should be evaluated when running a playbook or at least after Unpack().
	// Looks like this is the wrong place to presume that all child tasks are comming from child bobfiles
	// must exist.
	for i, task := range aggregate.BTasks {
		for _, dependentTaskName := range task.DependsOn {

			dependentTask, ok := aggregate.BTasks[dependentTaskName]
			if !ok {
				return nil, ErrTaskDoesNotExist
			}

			for exportname, export := range dependentTask.Exports {
				// fmt.Printf("Task %s exports %s\n", dependentTaskName, export)

				envvar := taskNameToEnvironment(dependentTaskName, exportname)

				value := filepath.Join(dependentTask.Dir(), string(export))

				// Make the path relative to the aggregates dir.
				dir := aggregate.Dir()
				if !strings.HasSuffix(dir, "/") {
					dir = dir + "/"
				}
				value = strings.TrimPrefix(value, dir)

				// println(envvar, value)

				task.AddEnvironment(envvar, value)

				aggregate.BTasks[i] = task
			}
		}
	}

	// Assure tasks are correctly initialised.
	for i, task := range aggregate.BTasks {
		task.WithLocalstore(b.local)
		task.WithBuildinfoStore(b.buildInfoStore)

		// a task must always-rebuild when caching is disabled
		if !b.enableCaching {
			task.SetRebuildStrategy(bobtask.RebuildAlways)
		}
		aggregate.BTasks[i] = task
	}

	return aggregate, aggregate.Verify()
}

func addTaskPrefix(prefix, taskname string) string {
	taskname = filepath.Join(prefix, taskname)
	taskname = strings.TrimPrefix(taskname, "/")
	return taskname
}

// taskNameToEnvironment
//
// Each taskname is translated into environment variables like:
//   `second-level/openapi_exportname => SECOND_LEVEL_OPENAPI_EXPORTNAME`
// Hyphens`-` are translated to underscores`_`.
func taskNameToEnvironment(taskname string, exportname string) string {

	splits := strings.Split(taskname, "/")
	splits = append(splits, exportname)

	envvar := strings.Join(splits, "_")
	envvar = strings.ReplaceAll(envvar, "-", "_")
	envvar = strings.ReplaceAll(envvar, ".", "_")
	envvar = strings.ToUpper(envvar)

	return envvar
}
