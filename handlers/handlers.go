package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Home returns the path of current request
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("Request catched %s.", r.URL.Path)
	fmt.Fprintf(w, "Processing URL %s...\n", r.URL.Path)
}
