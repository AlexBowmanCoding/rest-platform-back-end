package main

import (
	"log"
	"net/http"
	"github.com/AlexBowmanCoding/rest-platform-back-end/api/app"
	"github.com/rs/cors"
	
)




func main() {
	//initalize app
	app := app.App{}
	app.Initialize()

	//log output
	log.Print("Now Running Router")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"POST, GET, OPTIONS, PUT, DELETE"},
		AllowedHeaders: []string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"},
	})
	handler := c.Handler(app.Router)
	log.Fatal(http.ListenAndServe(":8001", handler))
	
}


