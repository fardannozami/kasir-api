package handlers

import (
	"kasir-api/utils"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.EncodeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
