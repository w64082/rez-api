package endpoints

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rez-api/app"
	"net/http"
)

type Visit struct {
	Id            string         `json:"id"`
	IdPlace       string         `json:"id_place"`
	IdWorker      string         `json:"id_worker"`
	DateStart     string         `json:"date_start"`
	DateTo        string         `json:"date_to"`
	IsReserved    bool           `json:"is_reserved"`
	ClientName    sql.NullString `json:"client_name"`
	ClientSurname sql.NullString `json:"client_surname"`
	CreatedAt     string         `json:"created_at"`
	DeletedAt     sql.NullString `json:"-"`
}

func AddVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	_ = r.ParseForm()

	idWorker := r.Form.Get("id_worker")
	idPlace := r.Form.Get("id_place")
	dateStart := r.Form.Get("date_start")
	dateTo := r.Form.Get("date_to")
	if idWorker == "" || idPlace == "" || dateStart == "" || dateTo == "" {
		response["message"] = "invalid visit input data"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `INSERT INTO visits (id_worker, id_place, date_start, date_to) VALUES ($1, $2, $3, $4) RETURNING id`
	lastInsertId := ""
	err := app.Container.DbHandle.QueryRow(sqlStatement, idWorker, idPlace, dateStart, dateTo).Scan(&lastInsertId)
	if err != nil {
		fmt.Print(err)
		response["message"] = "cannot add visit"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	response["message"] = "new visit added"
	response["id"] = lastInsertId
	w.WriteHeader(http.StatusCreated)
	jsonResp, _ := json.Marshal(response)
	w.Write(jsonResp)
	return
}

func GetAllVisits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	sqlStatement := `SELECT id, id_place, id_worker, date_start, date_to, is_reserved, client_name, client_surname, created_at, deleted_at FROM visits WHERE deleted_at IS NULL ORDER BY created_at DESC;`
	rows, _ := app.Container.DbHandle.Query(sqlStatement)

	defer rows.Close()
	var visits []Visit
	for rows.Next() {
		var v Visit
		if err := rows.Scan(&v.Id, &v.IdPlace, &v.IdWorker, &v.DateStart, &v.DateTo, &v.IsReserved, &v.ClientName, &v.ClientSurname, &v.CreatedAt, &v.DeletedAt); err != nil {
			response["message"] = "cannot fetch visits"
			w.WriteHeader(http.StatusInternalServerError)
			jsonResp, _ := json.Marshal(response)
			w.Write(jsonResp)
			return
		}
		visits = append(visits, v)
	}

	jsonResp, _ := json.Marshal(visits)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func GetVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid visit ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `SELECT id, id_place, id_worker, date_start, date_to, is_reserved, client_name, client_surname, created_at, deleted_at FROM visits WHERE id=$1 AND deleted_at IS NULL;`
	visit := Visit{}
	row := app.Container.DbHandle.QueryRow(sqlStatement, id)
	switch err := row.Scan(&visit.Id, &visit.IdPlace, &visit.IdWorker, &visit.DateStart, &visit.DateTo, &visit.IsReserved, &visit.ClientName, &visit.ClientSurname, &visit.CreatedAt, &visit.DeletedAt); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNoContent)
		return
	case nil:
		jsonResp, _ := json.Marshal(visit)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	default:
		response["message"] = "cannot fetch visit"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}
}

func VisitReservation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid visit ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	_ = r.ParseForm()
	inputName := r.Form.Get("client_name")
	inputSurname := r.Form.Get("client_surname")
	if inputName == "" || inputSurname == "" {
		response["message"] = "invalid client name or surname"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	app.Container.DbHandle.Query("UPDATE visits SET is_reserved = true, client_name = $1, client_surname = $2 WHERE id = $3", inputName, inputSurname, id)

	w.WriteHeader(http.StatusCreated)
	return
}

func DeleteVisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid visit ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	app.Container.DbHandle.Query("UPDATE visits SET deleted_at = NOW() WHERE id = $1", id)

	w.WriteHeader(http.StatusNoContent)
	return
}
