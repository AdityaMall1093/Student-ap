package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Adityamall1093/student-api/internal/storage"
	"github.com/Adityamall1093/student-api/internal/types"
	"github.com/Adityamall1093/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
		slog.Info("creaing a student")

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.ErrUnexpectedEOF) {

			response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("Empty body")))
			return
		}
		if err != nil {

			response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
			return
		}
		//request validation

		// w.Write([]byte("Welcome to Student API"))

		// if err := validator.New().Struct(student); err != nil {
		// 	validateErrs := err.(validator.ValidationErrors)
		// 	response.WriteJson(w, http.StatusBadRequest, response.ValidationError(err))
		// 	return
		// }

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)

			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		lastid, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("user created sucessfully", slog.String("userId", fmt.Sprint(lastid)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{
			"id": lastid,
		})

		// response.WriteJson(w, http.StatusCreated, map[string]string{"succes": "ok"})
	}
}
