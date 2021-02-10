package common

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(r *http.Request, w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithStatus(w http.ResponseWriter, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
}
