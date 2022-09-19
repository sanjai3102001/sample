package main

import (
	"log"
	"net/http"
	"testcalling/actualfunc"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	// routes.Routeuser(mainRouter)
	log.Fatal(http.ListenAndServe(":8080", mainRouter))

	actualfunc.Afunc()
}
