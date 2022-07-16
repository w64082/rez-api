package endpoints

import (
	"encoding/json"
	"net/http"
)

func GetAllVisits(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "get all visits"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func AddVisit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "add visity"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func GetVisit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "get visit"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func VisitReservation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "reservstion"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func DeleteVisit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "delete visit"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
