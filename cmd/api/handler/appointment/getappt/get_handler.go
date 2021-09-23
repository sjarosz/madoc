package getappt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func GetApptById(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		apptId := r.Context().Value(entity.CtxKey("apptId"))
		appt, err := app.AppointmentService.GetAppointment(r.Context(), apptId.(int))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "no appointments available")
			return
		}

		response, _ := json.Marshal(appt)
		w.Write(response)
	}
}

func GetAllAppointments(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		username := r.Context().Value(entity.CtxKey("username"))
		appts, err := app.AppointmentService.GetAllAppointments(r.Context(), username.(string))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		if len(*appts) > 0 {
			response, _ := json.Marshal(appts)
			w.Write(response)
		} else {
			fmt.Fprint(w, "No appointments available")
		}
	}
}

func HandleGetRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateGetRequest,
	}

	return middleware.Chain(GetApptById(app), app, mdw...)
}

func HandleGetAllRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateGetAllRequest,
	}

	return middleware.Chain(GetAllAppointments(app), app, mdw...)
}
