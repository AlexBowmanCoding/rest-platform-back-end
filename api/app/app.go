package app

import (
	"github.com/AlexBowmanCoding/content-hub-back-end/mongoDB"
	"github.com/AlexBowmanCoding/content-hub-back-end/user"
	"github.com/gorilla/mux"
)

type App struct{
	Router *mux.Router
	DB mongodb.MongoDB
	UserMethods user.Mongo
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.DB = mongodb.NewMongoDB()
	app.UserMethods = user.Mongo{
		DB: app.DB,
		Client: app.DB.Client,
	}


	app.Router.HandleFunc("/users", app.UserMethods.NewUser).Methods("POST")
	app.Router.HandleFunc("/users/login/{id}", app.UserMethods.LoginUser).Methods("GET")
	app.Router.HandleFunc("/users/{id}", app.UserMethods.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id}", app.UserMethods.DeleteUser).Methods("DELETE")
	
}