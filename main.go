package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/devfest-2017/handlers"
)

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/any
func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set.")
	}

	log.Print("Initialize service...")

	router := httprouter.New()
	router.GET("/", handlers.Home)

	router.GET("/version", handlers.Version)

	log.Print("We are ready to handle requests.")
	log.Print(http.ListenAndServe(":"+port, router))
}
