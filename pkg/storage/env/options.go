package env

type options struct {
	Prefix string
}

type option func(*options)

func WithPrefix(val string) option {
	return func(args *options) {
		args.Prefix = val
	}
}
