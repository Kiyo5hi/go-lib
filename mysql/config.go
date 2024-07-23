package mysql

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	Username string
	Password string
	Protocol string
	Host     string
	Port     uint16
	Database string
	options  map[string]string
}

type Option func(c *Config) error

func NewConfig(username, password, protocol, host string, port uint16, database string, opts ...Option) (*Config, error) {
	conf := &Config{
		Username: username,
		Password: password,
		Protocol: protocol,
		Host:     host,
		Port:     port,
		Database: database,
	}

	for _, opt := range opts {
		if err := opt(conf); err != nil {
			return nil, fmt.Errorf("failed to create mysql config: %w", err)
		}
	}

	return conf, nil
}

func WithCharset(charset string) Option {
	return func(c *Config) error {
		c.options["charset"] = charset
		return nil
	}
}

func WithPraseTime(parseTime bool) Option {
	pt := "False"
	if parseTime {
		pt = "True"
	}
	return func(c *Config) error {
		c.options["parseTime"] = pt
		return nil
	}
}

func WithTimezone(location *time.Location) Option {
	return func(c *Config) error {
		c.options["loc"] = location.String()
		return nil
	}
}

// Dsn: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
// Example: root:root@tcp(mysql-dev:3306)/gmwe?charset=utf8mb4&parseTime=True&loc=Local
func (c *Config) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", c.Username, c.Password, c.Protocol, c.Host, c.Port, c.Database)
	if len(c.options) > 0 {
		options := []string{}
		for k, v := range c.options {
			options = append(options, fmt.Sprintf("%s=%s", k, v))
		}
		query := strings.Join(options, "&")
		query = url.QueryEscape(query)
		dsn = fmt.Sprintf("%s?%s", dsn, query)
	}
	return dsn
}
