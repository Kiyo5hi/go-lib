package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	Username string
	Password string
	Database string
	Options  *options.ClientOptions
}

type Option func(c *Config) error

func NewConfig(host, username, password, database string, opts ...Option) (*Config, error) {
	config := &Config{
		Host:     host,
		Username: username,
		Password: password,
		Database: database,
	}

	for _, opt := range opts {
		err := opt(config)
		if err != nil {
			return nil, fmt.Errorf("failed to create mongo config: %w", err)
		}
	}

	return config, nil
}

func WithOptions(opts *options.ClientOptions) Option {
	return func(c *Config) error {
		c.Options = opts
		return nil
	}
}

// Srv genreates Mongo URI using the config passed in.
// format: mongodb+srv://[username:password@]host[/[defaultauthdb][?options]]
func (c *Config) Srv() string {
	srv := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", c.Username, c.Password, c.Host, c.Database)
	return srv
}
