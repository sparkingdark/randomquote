package server

import (
	"context"
	// "fmt"
	"log"
	// "time"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
)


func connect() *mongo.Client{
	credential := options.Credential{
		Username: "root",
		Password: "example",
	}
	client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func Insertdata(collection *mongo.Collection, data interface{}) error {
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func Deletedata(collection *mongo.Collection, data interface{}) error {
	_, err := collection.DeleteOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func GetData(collection *mongo.Collection, data interface{}) error {
	err := collection.FindOne(context.TODO(), data).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func Casewise(data interface{},command string) {
	client := connect()
	// ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	collection := client.Database("randoquote").Collection("quotes")
	switch command {
		case "insert":
			Insertdata(collection, data)
		case "delete":
			Deletedata(collection, data)
		case "get":
			GetData(collection, data)
		default:
			log.Fatal("Invalid type")
	}
}