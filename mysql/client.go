package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	InnerClient *gorm.DB
	Config      *Config
}

func NewClient(conf *Config) (*Client, error) {
	dsn := conf.Dsn()
	client, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to create mysql client: %w", err)
	}
	return &Client{
		InnerClient: client,
		Config:      conf,
	}, nil
}

func (c *Client) AutoMigrate(dals ...any) error {
	err := c.InnerClient.AutoMigrate(dals)
	if err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}
	return nil
}
