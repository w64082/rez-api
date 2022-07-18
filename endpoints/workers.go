package endpoints

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rez-api/app"
	"net/http"
)

type Worker struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	CreatedAt string         `json:"created_at"`
	DeletedAt sql.NullString `json:"-"`
}

func AddWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	_ = r.ParseForm()

	inputName := r.Form.Get("name")
	inputSurname := r.Form.Get("surname")
	if inputName == "" || inputSurname == "" {
		response["message"] = "invalid worker name or surname"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `INSERT INTO workers (name, surname) VALUES ($1, $2) RETURNING id`
	lastInsertId := ""
	err := app.Container.DbHandle.QueryRow(sqlStatement, inputName, inputSurname).Scan(&lastInsertId)
	if err != nil {
		response["message"] = "cannot add worker"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	response["message"] = "new worker added"
	response["id"] = lastInsertId
	w.WriteHeader(http.StatusCreated)
	jsonResp, _ := json.Marshal(response)
	w.Write(jsonResp)
	return
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid worker ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `SELECT id, name, surname, created_at, deleted_at FROM workers WHERE id=$1 AND deleted_at IS NULL;`
	worker := Worker{}
	row := app.Container.DbHandle.QueryRow(sqlStatement, id)
	switch err := row.Scan(&worker.Id, &worker.Name, &worker.Surname, &worker.CreatedAt, &worker.DeletedAt); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNoContent)
		return
	case nil:
		jsonResp, _ := json.Marshal(worker)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	default:
		response["message"] = "cannot fetch worker"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}
}

func GetAllWorkers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	sqlStatement := `SELECT id, name, surname, created_at, deleted_at FROM workers WHERE deleted_at IS NULL ORDER BY created_at DESC;`
	rows, _ := app.Container.DbHandle.Query(sqlStatement)

	defer rows.Close()
	var workers []Worker
	for rows.Next() {
		var wk Worker
		if err := rows.Scan(&wk.Id, &wk.Name, &wk.Surname, &wk.CreatedAt, &wk.DeletedAt); err != nil {
			response["message"] = "cannot fetch workers"
			w.WriteHeader(http.StatusInternalServerError)
			jsonResp, _ := json.Marshal(response)
			w.Write(jsonResp)
			return
		}
		workers = append(workers, wk)
	}

	jsonResp, _ := json.Marshal(workers)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid worker ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	app.Container.DbHandle.Query("UPDATE workers SET deleted_at = NOW() WHERE id = $1", id)

	w.WriteHeader(http.StatusNoContent)
	return
}
