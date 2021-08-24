package createappt

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
	"github.com/sqoopdata/madoc/pkg/middleware"
)

func create(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		aO := r.Context().Value(model.CtxKey("appt"))
		appt := aO.(*model.Appointment)

		if err := appt.Create(r.Context(), app); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		response, _ := json.Marshal(appt)
		w.Write(response)
	}
}

func HandleRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequest,
	}

	return middleware.Chain(create(app), app, mdw...)
}
