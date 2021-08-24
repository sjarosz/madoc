package getappt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
	"github.com/sqoopdata/madoc/pkg/middleware"
)

func GetApptByApptId(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		apptId := r.Context().Value(model.CtxKey("apptId"))
		appt := &model.Appointment{ApptId: apptId.(int)}

		if err := appt.GetByApptId(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "no appointments available")
			return
		}

		response, _ := json.Marshal(appt)
		w.Write(response)
	}
}

func GetApptByUsername(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		username := r.Context().Value(model.CtxKey("username"))
		appt := &model.Appointment{Patient: username.(string)}

		appts, err := appt.GetByUsername(r.Context(), app)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		if len(appts) > 0 {
			response, _ := json.Marshal(appts)
			w.Write(response)
		} else {
			fmt.Fprintf(w, "No appointments available")
		}
	}
}

func HandleRequestByUsername(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequest,
	}

	return middleware.Chain(GetApptByUsername(app), app, mdw...)
}

func HandleRequestByApptId(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequestById,
	}

	return middleware.Chain(GetApptByApptId(app), app, mdw...)
}
