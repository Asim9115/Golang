package controller

import (
	"context"
	"log"
	"os"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString = os.Getenv("MONGO_CONNECTION_STRING")
var dbName = "netflix"
var colName = "watchlist"

//Most important 
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	//connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption )

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).collection(colName)


}