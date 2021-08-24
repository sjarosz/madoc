package deleteappt

import (
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
	"github.com/sqoopdata/madoc/pkg/middleware"
)

func DeleteApptByApptId(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		apptId := r.Context().Value(model.CtxKey("apptId"))
		appt := &model.Appointment{ApptId: apptId.(int)}

		if err := appt.DeleteByApptId(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		fmt.Fprintf(w, "Appointment Deleted")
	}
}

func HandleRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequestById,
	}

	return middleware.Chain(DeleteApptByApptId(app), app, mdw...)
}
