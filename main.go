package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/xizahk/gowebsite/app/controller"
)

func newRouter() *mux.Router {
	// Create a new mux Router
	r := mux.NewRouter()

	// Register Handlers and URL paths
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("", http.FileServer(staticFileDirectory))
	r.HandleFunc("/userWithImages", controller.GetUsersWithImagesHandler).Methods("GET")
	r.PathPrefix("").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	log.Println("Starting web server...")

	// Calls newRouter() to create a new mux router and host the website at http://localhost:8080/
	r := newRouter()
	log.Println("Serving web server at http://localhost:8080/")
	http.ListenAndServe(":8080", r)
}
