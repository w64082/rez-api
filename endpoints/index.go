package endpoints

import (
	"encoding/json"
	"net/http"
)

func ApiIndexPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "API is ready to use"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
