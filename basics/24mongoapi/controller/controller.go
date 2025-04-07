package controller

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//const mongoConnectionString = "mongodb+srv://ashwanibhadani3289:<db_password>@notetakercluster.inmx7.mongodb.net/?retryWrites=true&w=majority&appName=noteTakerCluster"

func ConnectToMongoAtlas() *mongo.Client {
	//LOAD THE .ENV FILE
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading up credentials from .env file!")
	}

	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	cluster := os.Getenv("MONGO_CLUSTER")

	//build the MongoDB URI
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, pass, cluster)

	//connect to mongodb, setup timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//apply connection options
	opts := options.Client().ApplyURI(uri)
	//connect to MongoDB
	client, err := mongo.Connect(opts)

	if err != nil {
		log.Fatalf("Falied to connect : %v\n", err)
	}

	//ping MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Could not ping MongoDB:%v\n", err)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	return client
}
