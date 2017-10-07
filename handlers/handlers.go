package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rumyantseva/devfest-2017/version"
)

// Home returns the path of current request
func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("Request catched %s.", r.URL.Path)
	fmt.Fprintf(w, "Processing URL %s...\n", r.URL.Path)
}

// Version is a handler to show current version
func Version(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Version us %s, Repo is %s, Commit is %s", version.RELEASE, version.REPO, version.COMMIT)
}
