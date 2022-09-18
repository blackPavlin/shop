package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDB
func NewMongoDB(ctx context.Context, databaseURL string, databaseName string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL).SetDirect(true))
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(databaseName), nil
}
