package mongodb

import (
	"context"
	"github.com/jvhab/Go-clean/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

func NewMongoClient(ctx context.Context, cfg *config.Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Mongodb.URI).SetAuth(options.Credential{
		Username: cfg.Mongodb.User,
		Password: cfg.Mongodb.Password,
	}).SetConnectTimeout(connectTimeout).SetMaxConnIdleTime(maxConnIdleTime).SetMinPoolSize(minPoolSize).SetMaxPoolSize(maxPoolSize))
	if err != nil {
		return nil, err
	}
	if err = client.Connect(ctx); err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}
