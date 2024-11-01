package logger

type options struct {
	appName string
	debug   bool
}

type option func(*options)

func Options(opts ...option) *options {
	defaultOptions := &options{
		debug:   false,
		appName: "app",
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

func WithAppName(appName string) option {
	return func(opts *options) {
		opts.appName = appName
	}
}
