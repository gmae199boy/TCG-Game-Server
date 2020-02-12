package dbserver

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//ChaUpdate : 캐릭터를 추가하거나 변경할 떄 쓰는 구조체
type ChaUpdate struct {
}

//ChaInsertMany : 새로운 캐릭터가 여러개 추가될 때
func (cu *ChaUpdate) ChaInsertMany(client *mongo.Client, insertNum int) {
	convertCard := make([]interface{}, insertNum)
	for i := 0; i < insertNum; i++ {
		//convertCard[i] = chaArr[i]
	}
	mod := mongo.IndexModel{
		Keys:    bsonx.Doc{{"id", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	collection := client.Database("test").Collection("chacard")
	ind, err := collection.Indexes().CreateOne(context.TODO(), mod)
	fmt.Println("createOne() ", ind)
	//인서트
	insertResult, err := collection.InsertMany(context.TODO(), convertCard)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedIDs)
}

//ChaInsert : 새로운 캐릭터를 1개 추가 할 때
func (cu *ChaUpdate) ChaInsert(client *mongo.Client) {
	mod := mongo.IndexModel{
		Keys:    bsonx.Doc{{"id", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}

	collection := client.Database("test").Collection("chacard")
	ind, err := collection.Indexes().CreateOne(context.TODO(), mod)
	fmt.Println("createOne() ", ind)
	//인서트
	//insertResult, err := collection.InsertOne(context.TODO(), chaCard)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
