package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGO_URI = "mongodb://docker:mongopw@localhost:49153"
)
var (
	Db = DBInstance()
	UsersColl = OpenCollection("users")
	AuthsColl = OpenCollection("auths")
	StoresColl = OpenCollection("stores")
	EmailColl = OpenCollection("emails")
)
func DBInstance() *mongo.Database{
	client,err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil{
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	return client.Database("gmail")
}
func OpenCollection(CollName string) *mongo.Collection{
	return Db.Collection(CollName)
}