package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rumyantseva/test2doc/example/foos"
	"github.com/rumyantseva/test2doc/example/widgets"
)

func main() {
	router := mux.NewRouter()
	foos.AddRoutes(router)
	widgets.AddRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
