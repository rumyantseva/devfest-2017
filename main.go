package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/devfest-2017/handlers"
)

// Run server: go build -o app && ./app
// Try requests: curl http://127.0.0.1:8000/any
func main() {
	router := httprouter.New()
	router.GET("/", handlers.Home)

	http.ListenAndServe(":8000", router)
}
