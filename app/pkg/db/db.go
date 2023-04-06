package db

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr    = Getenv("MONGO_USERNAME", "root")
	pwd    = Getenv("MONGO_PASSWORD", "secret")
	host   = Getenv("MONGO_HOST", "mongo")
	port   = Getenv("MONGO_PORT", "27017")
	dbname = Getenv("MONGO_DATABASE", "mongo")
)

func GetCollection() *mongo.Collection {
	int_port, err := strconv.Atoi(port)
	if err != nil {
		panic(err.Error())
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, int_port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}
	return client.Database(dbname).Collection("Destiny")
}
