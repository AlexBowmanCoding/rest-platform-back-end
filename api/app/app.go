package app

import (
	"main/mongoDB"
	"main/user"

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
	app.Router.HandleFunc("/users/login/{id}", user.LoginUser).Methods("POST")
	app.Router.HandleFunc("/users/login/{id}", user.DeleteUser).Methods("DELETE")
	
}