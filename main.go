package main

import (
	"fmt"
	"github.com/rez-api/app"
	"github.com/rez-api/endpoints"
	"github.com/rez-api/resources"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", endpoints.ApiIndexPage)

	container := &app.Container
	container.DbHandle = resources.GetPostgresConnection()

	// routes for places
	router.HandleFunc("/places", endpoints.AddPlace).Methods("POST")                          // add new
	router.HandleFunc("/places", endpoints.GetAllPlaces).Methods("GET")                       // get all
	router.HandleFunc("/places/{id:[0-9a-z_-]{36}}", endpoints.GetPlace).Methods("GET")       // get single by id
	router.HandleFunc("/places/{id:[0-9a-z_-]{36}}", endpoints.DeletePlace).Methods("DELETE") // delete

	// routes for workers
	router.HandleFunc("/workers", endpoints.AddWorker).Methods("POST")                          // add new
	router.HandleFunc("/workers", endpoints.GetAllWorkers).Methods("GET")                       // get all
	router.HandleFunc("/workers/{id:[0-9a-z_-]{36}}", endpoints.GetWorker).Methods("GET")       // get single by id
	router.HandleFunc("/workers/{id:[0-9a-z_-]{36}}", endpoints.DeleteWorker).Methods("DELETE") // delete

	// routes for visits
	router.HandleFunc("/visits", endpoints.AddVisit).Methods("POST")                                        // add new
	router.HandleFunc("/visits", endpoints.GetAllVisits).Methods("GET")                                     // get all
	router.HandleFunc("/visits/{id:[0-9a-z_-]{36}}", endpoints.GetVisit).Methods("GET")                     // get single by id
	router.HandleFunc("/visits/{id:[0-9a-z_-]{36}}/reservation", endpoints.VisitReservation).Methods("PUT") // make reservation for visit
	router.HandleFunc("/visits/{id:[0-9a-z_-]{36}}", endpoints.DeleteVisit).Methods("DELETE")               // delete

	fmt.Println("Starting API")
	log.Fatal(http.ListenAndServe(":8080", router))
}
