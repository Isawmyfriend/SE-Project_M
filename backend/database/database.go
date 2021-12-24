package database

import (
	"backend/config"
	_ "backend/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func DB() (*mongo.Client,context.Context) {

	conf := config.GetConfig()
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.GetString("database.host")))
    if err != nil {
        log.Fatal(err)
    }

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	return client,ctx

	
}

