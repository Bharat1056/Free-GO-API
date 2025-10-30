package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Bharat1056/students-api/internal/response"
	"github.com/Bharat1056/students-api/internal/types"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		// if body is empty then
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			// response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // custom error message
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) // typecast the error to the validation error types
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
		}

		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string {"success" : "ok"})
	}
}
