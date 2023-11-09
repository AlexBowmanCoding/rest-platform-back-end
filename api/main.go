package main

import (
	"log"
	"net/http"
	"github.com/AlexBowmanCoding/content-hub-back-end/app"

	
)




func main() {
	//initalize app
	app := app.App{}
	app.Initialize()

	//log output
	log.Print("Now Running Mux Router")
	log.Fatal(http.ListenAndServe(":8001", app.Router))
	
}


