package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	URI    string
}

func NewMongoDB() MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb+srv://contenthub.lgtdzvq.mongodb.net/").SetAuth(options.Credential{
		Username: "Admin",
		Password: "Admin", //CHANGE LATER WHEN MAKING AUTH
	})
	newClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = newClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return MongoDB{
		Client: newClient,
		URI:    "mongodb+srv://contenthub.lgtdzvq.mongodb.net/",
	}
}

func (db MongoDB) Get() {

}

func (db MongoDB) Post(c mongo.Collection, item interface{}) error {
	_, err := c.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}
	return nil
}

func (db MongoDB) Update() {

}

func (db MongoDB) Delete() {

}
