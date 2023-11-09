package user

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB holds all the variables and methods associated with the Database
type MongoDB struct {
	Client *mongo.Client
	URI    string
}



// NewMongoDB creates a new MongoDB struct and connects to existing database
func NewUserDB() MongoDB {
	clientOptions := options.Client().ApplyURI("mongodb+srv://contenthub.lgtdzvq.mongodb.net/").SetAuth(options.Credential{
		Username: "Admin",
		Password: "Admin", //Change When making authentication
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

// Get returns a mongo item from the Database
func (db MongoDB) Get(c mongo.Collection, id string) (User, error) {
	filter := bson.D{{"id", id}}
	var result User
	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Post puts an item in the Database
func (db MongoDB) Post(c mongo.Collection, user User) error {
	_, err := c.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// Update Updates an item in the database
func (db MongoDB) Update(c mongo.Collection, user User) error {
	filter := bson.D{{"id", user.ID}}

	update := bson.D{
		{"$set", bson.D{
			{"username", user.Username},
			{"password", user.Password},
			},
		},
	}
	_, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
// Delete Deletes an item in the database.
func (db MongoDB) Delete(c mongo.Collection, id string) error {
	_, err := c.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		return err
	}
	return nil
}
