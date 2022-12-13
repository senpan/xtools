package confx

type (
	// Option defines the method to customize the config options.
	Option func(opt *options)

	options struct {
		env        bool
		config     string
		pathPrefix string
	}
)

// WithEnv customizes the config to use environment variables.
func WithEnv() Option {
	return func(opt *options) {
		opt.env = true
	}
}

// WithConfig customizes the config file
func WithConfig(path string) Option {
	return func(opt *options) {
		opt.config = path
	}
}

// WithConfigPathPrefix customizes the config prefix path
func WithConfigPathPrefix(path string) Option {
	return func(opt *options) {
		opt.pathPrefix = path
	}
}
