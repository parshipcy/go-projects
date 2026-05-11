package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Parship12/bookstore-api/pkg/routes"
)

func main(){
	r := mux.NewRouter() // creates a new mux router instance
	routes.RegisterBookStoreRoutes(r) // register bookstore routes on the router
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // Start the server and use that router to handle all requests
}

/*
The problem:
When you pass a handler (like r) directly to ListenAndServe, 
it uses that handler and ignores the default mux. So http.Handle had no effect.
*/