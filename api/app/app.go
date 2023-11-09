package app

import (
	"github.com/AlexBowmanCoding/content-hub-back-end/user"
	"github.com/gorilla/mux"
)

type App struct {

	// Router handles http routing
	Router *mux.Router

	// DB holds all the variables and methods associated with the Database
	DB          user.MongoDB

	//User holds the mongo client and DB variables along with the user methods 
	User user.MongoUser
}
// Initialize creates a new mux router and connects to the mongo database 
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.DB = user.NewUserDB()
	app.User = user.MongoUser{
		DB:     app.DB,
		Client: app.DB.Client,
	}

	app.Router.HandleFunc("/users", app.User.NewUser).Methods("POST")
	app.Router.HandleFunc("/users/login/{id}", app.User.LoginUser).Methods("GET")
	app.Router.HandleFunc("/users/{id}", app.User.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id}", app.User.DeleteUser).Methods("DELETE")
	app.Router.HandleFunc("/users/{id}", app.User.GetUser).Methods("GET")
}
