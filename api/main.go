package main

import (
	"log"
	"net/http"
	"main/app"

	
)




func main() {

	app := app.App{}
	app.Initialize()

	log.Print("Now Running Mux Router")
	log.Fatal(http.ListenAndServe(":8001", app.Router))
	
}


