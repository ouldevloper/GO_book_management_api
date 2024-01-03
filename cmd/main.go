package main

import (
	"log"
	"net/http"
	"restapi/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRouter(router)

	log.Fatal("")
	http.ListenAndServe(":1234", nil)
}
