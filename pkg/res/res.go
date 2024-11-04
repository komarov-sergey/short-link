package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
