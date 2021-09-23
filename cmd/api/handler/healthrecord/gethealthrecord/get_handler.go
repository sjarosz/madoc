package gethealthrecord

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func getHealthRecordByPatient(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		patient := r.Context().Value(entity.CtxKey("username"))

		records, err := app.HealthRecordService.GetByPatient(r.Context(), patient.(string))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		if len(*records) > 0 {
			response, _ := json.Marshal(records)
			w.Write(response)
		} else {
			fmt.Fprint(w, "No records found!")
		}
	}
}

func HandleGetByPatientRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateGetRequest,
	}

	return middleware.Chain(getHealthRecordByPatient(app), app, mdw...)
}
