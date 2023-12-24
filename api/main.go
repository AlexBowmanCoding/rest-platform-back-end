package main

import (
	"log"
	"net/http"
	"github.com/AlexBowmanCoding/content-hub-back-end/app"
	"github.com/rs/cors"
	
)




func main() {
	//initalize app
	app := app.App{}
	app.Initialize()

	//log output
	log.Print("Now Running Mux Router")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:3000"},
		AllowCredentials: true,
	})
	handler := c.Handler(app.Router)
	log.Fatal(http.ListenAndServe(":8001", handler))
	
}


