package config

import (
	"context"
	"dbclient/constants"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	mongoconnection := options.Client().ApplyURI(constants.ConnectionStrings)

	mongoclient, err := mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err

	}
	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	fmt.Println("DataBase Connected")
	return mongoclient, nil

}
