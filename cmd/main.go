package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRouter(router)
	http.Handle("/", router)
	args := os.Args

	fmt.Println("Server Started On : http://127.0.0.1:" + args[1])
	log.Fatal(http.ListenAndServe(":"+args[1], nil))
}
