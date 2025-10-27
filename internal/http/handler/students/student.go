package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Bharat1056/students-api/internal/response"
	"github.com/Bharat1056/students-api/internal/types"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		// if body is empty then
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string {"success" : "ok"})
	}
}
