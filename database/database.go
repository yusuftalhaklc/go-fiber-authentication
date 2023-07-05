package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yusuftalhaklc/go-fiber-authentication/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

// Connect establishes a connection to the MongoDB database.
func Connect() {
	connectionString := config.Config("DB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	Db = client.Database(config.Config("DB_NAME"))

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	log.Println("MongoDB Connection established!")
}
