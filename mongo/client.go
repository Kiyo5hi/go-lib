package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	InnerClient *mongo.Client
	Config      *Config
}

func NewClient(ctx context.Context, conf *Config) (*Client, error) {
	srv := conf.Srv()
	opts := conf.Options
	if opts == nil {
		options.Client()
	}
	opts.ApplyURI(srv)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create mongo client: %w", err)
	}
	return &Client{
		InnerClient: client,
		Config:      conf,
	}, nil
}

func (c *Client) Database(opts ...*options.DatabaseOptions) *mongo.Database {
	return c.InnerClient.Database(c.Config.Database, opts...)
}

func (c *Client) Collection(coll string, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.InnerClient.Database(c.Config.Database).Collection(coll, opts...)
}
