package main

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	mClient *mongo.Client
	ctx, _  = context.WithTimeout(context.Background(), 10*time.Second)
)

func connectMongo() {

	var err error
	mClient, err = mongo.Connect(ctx, mCfg.uri)

	if err != nil {
		log.Fatal("Could not connect to mongodb: ", err.Error())
	}

}

func getCollection(c string) *mongo.Collection {
	return mClient.Database(mCfg.db).Collection(c)
}
