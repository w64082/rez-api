package endpoints

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rez-api/app"
	"net/http"
)

type Place struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	DeletedAt sql.NullString `json:"-"`
}

func AddPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)
	_ = r.ParseForm()

	inputName := r.Form.Get("name")
	if inputName == "" {
		response["message"] = "invalid place name"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `INSERT INTO places (name) VALUES ($1) RETURNING id`
	lastInsertId := ""
	err := app.Container.DbHandle.QueryRow(sqlStatement, inputName).Scan(&lastInsertId)
	if err != nil {
		response["message"] = "cannot add place"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	response["message"] = "new place added"
	response["id"] = lastInsertId
	w.WriteHeader(http.StatusCreated)
	jsonResp, _ := json.Marshal(response)
	w.Write(jsonResp)
	return
}

func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	sqlStatement := `SELECT id, name, created_at, deleted_at FROM places WHERE deleted_at IS NULL ORDER BY created_at DESC;`
	rows, _ := app.Container.DbHandle.Query(sqlStatement)

	defer rows.Close()
	var places []Place
	for rows.Next() {
		var p Place
		if err := rows.Scan(&p.Id, &p.Name, &p.CreatedAt, &p.DeletedAt); err != nil {
			response["message"] = "cannot fetch places"
			w.WriteHeader(http.StatusInternalServerError)
			jsonResp, _ := json.Marshal(response)
			w.Write(jsonResp)
			return
		}
		places = append(places, p)
	}

	jsonResp, _ := json.Marshal(places)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func GetPlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid place ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	sqlStatement := `SELECT id, name, created_at, deleted_at FROM places WHERE id=$1 AND deleted_at IS NULL;`
	place := Place{}
	row := app.Container.DbHandle.QueryRow(sqlStatement, id)
	switch err := row.Scan(&place.Id, &place.Name, &place.CreatedAt, &place.DeletedAt); err {
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNoContent)
		return
	case nil:
		jsonResp, _ := json.Marshal(place)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
		return
	default:
		response["message"] = "cannot fetch place"
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}
}

func DeletePlace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := make(map[string]string)

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		response["message"] = "invalid place ID"
		w.WriteHeader(http.StatusBadRequest)
		jsonResp, _ := json.Marshal(response)
		w.Write(jsonResp)
		return
	}

	app.Container.DbHandle.Query("UPDATE places SET deleted_at = NOW() WHERE id = $1", id)

	w.WriteHeader(http.StatusNoContent)
	return
}
