package gethealthrecord

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
	"github.com/sqoopdata/madoc/pkg/middleware"
)

func GetHealthRecordByRecordId(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		recordId := r.Context().Value(model.CtxKey("recordId"))
		record := &model.HealthRecord{HealthRecordId: recordId.(int)}

		if err := record.GetByHealthRecordId(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		response, _ := json.Marshal(record)
		w.Write(response)
	}
}

func GetHealthRecordByUsername(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		username := r.Context().Value(model.CtxKey("username"))
		record := &model.HealthRecord{Patient: username.(string)}

		records, err := record.GetByUsername(r.Context(), app)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		if len(records) > 0 {
			response, _ := json.Marshal(records)
			w.Write(response)
		} else {
			fmt.Fprintf(w, "No health records available")
		}
	}
}

func HandleRequestByUsername(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequest,
	}

	return middleware.Chain(GetHealthRecordByUsername(app), app, mdw...)
}

func HandleRequestByHealthRecordId(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequestById,
	}

	return middleware.Chain(GetHealthRecordByRecordId(app), app, mdw...)
}
