package dbserver

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB : DB조각 구조체
type DB struct {
}

//ConnectDB : DB에 연결한다.
func (d *DB) ConnectDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	// var card []ChaCard = make([]ChaCard, 5)
	// card[0].ID = 1
	// card[1].ID = 2
	// card[2].ID = 1
	// var cu ChaUpdate
	// cu.ChaInsertMany(client, 3, card)
}
