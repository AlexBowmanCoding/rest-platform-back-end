package mongo

import(
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB holds all the variables and methods associated with the Database
type MongoDB struct {
	Client *mongo.Client
	URI    string
}



// NewMongoDB creates a new MongoDB struct and connects to existing database
func NewDB() MongoDB {
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