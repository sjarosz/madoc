package updatehealthrecord

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func update(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		rO := r.Context().Value(entity.CtxKey("record"))
		record := rO.(*entity.HealthRecord)

		cRecord, err := app.HealthRecordService.UpdateHealthRecord(r.Context(), record)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		response, _ := json.Marshal(cRecord)
		w.Write(response)
	}
}

func HandleUpdateRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateUpdateRequest,
	}

	return middleware.Chain(update(app), app, mdw...)
}
