package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	m "github.com/AlexBowmanCoding/content-hub-back-end/mongo"
)

// User struct for holding user data. 
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// MongoUser struct for holding the mongo client and the mongoDB struct.
type MongoUser struct {
	Client *mongo.Client
	DB    m.MongoDB
}

// Get returns a mongo item from the Database
func (db MongoUser) Get(c mongo.Collection, id string) (User, error) {
	filter := bson.D{{"id", id}}
	var result User
	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Post puts an item in the Database
func (db MongoUser) Post(c mongo.Collection, user User) error {
	_, err := c.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// Update Updates an item in the database
func (db MongoUser) Update(c mongo.Collection, user User) error {
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
func (db MongoUser) Delete(c mongo.Collection, id string) error {
	_, err := c.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		return err
	}
	return nil
}
