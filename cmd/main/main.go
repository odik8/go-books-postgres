package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/odik8/go-books/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
	
}