package main

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "API is ready to use"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	// routes for places
	router.HandleFunc("/places", createEvent).Methods("POST") // add new
	router.HandleFunc("/places", createEvent).Methods("GET") // get all
	router.HandleFunc("/places/{id}", createEvent).Methods("GET") // get single by id
	router.HandleFunc("/places/{id}", createEvent).Methods("DELETE") // delete


	// routes for workers
	router.HandleFunc("/workers", createEvent).Methods("POST") // add new
	router.HandleFunc("/workers", createEvent).Methods("GET") // get all
	router.HandleFunc("/workers/{id}", createEvent).Methods("GET") // get single by id
	router.HandleFunc("/workers", createEvent).Methods("DELETE") // delete

	// routes for visits
	router.HandleFunc("/visits", createEvent).Methods("POST") // add new
	router.HandleFunc("/visits", createEvent).Methods("GET") // get all
	router.HandleFunc("/visits/{id}", createEvent).Methods("GET") // get single by id
	router.HandleFunc("/visits", createEvent).Methods("DELETE") // delete

	// routes for reservation
	router.HandleFunc("/reservation/{id}", createEvent).Methods("PUT") // add new - make reservation

	fmt.Println("Starting API");
	log.Fatal(http.ListenAndServe(":8080", router))
}