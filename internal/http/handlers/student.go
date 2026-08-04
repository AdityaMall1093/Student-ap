package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Adityamall1093/student-api/internal/types"
	"github.com/Adityamall1093/student-api/internal/utils/response"
)

func New() http.HandlerFunc {
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

		w.Write([]byte("Welcome to Student API"))

		response.WriteJson(w, http.StatusCreated, map[string]string{"succes": "ok"})
	}
}
