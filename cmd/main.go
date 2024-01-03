package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRouter(router)
	http.Handle("/", router)
	fmt.Println("Server Started On : http://127.0.0.1:1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
