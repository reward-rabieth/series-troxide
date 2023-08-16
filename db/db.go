package db

import (
	"context"
	"github.com/reward-rabieth/series-troxide/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect(cfg config.DatabaseConfig) (*mongo.Database, error) {
	log.Println("connecting to mongodb on address:  " + cfg.URI())
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.URI()))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Info("connected to mongodb on address", cfg.URI())
	return client.Database(cfg.DbName), nil

}
