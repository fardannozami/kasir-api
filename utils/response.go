package utils

import (
	"encoding/json"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func DecodeJSON(r *http.Request, data any, w http.ResponseWriter) bool {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		EncodeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return false
	}
	return true
}
