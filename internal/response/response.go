package response

import (
	"encoding/json"
	"net/http"
)


type Response struct {
	Status string
	Error string
}

// struct -> json --> encode
// json -> struct --> decode

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	
}
