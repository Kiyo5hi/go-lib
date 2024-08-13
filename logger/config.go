package logger

type options struct {
	debug bool
}

type option func(*options)

func Options(opts ...option) *options {
	defaultOptions := &options{
		debug: false,
	}
	for _, opt := range opts {
		opt(defaultOptions)
	}
	return defaultOptions
}

func WithDebug(debug bool) option {
	return func(opts *options) {
		opts.debug = debug
	}
}
