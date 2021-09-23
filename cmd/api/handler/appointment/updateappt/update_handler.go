package updateappt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func UpdateAppt(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		aObj := r.Context().Value(entity.CtxKey("appt"))
		appt := aObj.(*entity.Appointment)

		cAppt, err := app.AppointmentService.UpdateAppointment(r.Context(), appt)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Oops")
			return
		}

		response, _ := json.Marshal(cAppt)
		w.Write(response)
	}
}

func HandleUpdateRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateUpdateRequest,
	}

	return middleware.Chain(UpdateAppt(app), app, mdw...)
}
