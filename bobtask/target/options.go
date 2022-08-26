package target

type Option func(t *T)

func WithType(typee Type) Option {
	return func(t *T) {
		t.TypeSerialize = typee
	}
}

func WithTargetPaths(targetPaths []string) Option {
	return func(t *T) {
		t.PathsSerialize = targetPaths
	}
}
