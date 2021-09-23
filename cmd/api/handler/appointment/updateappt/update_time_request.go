package updateappt

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/appointmentvalidation"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func validateUpdateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appointment, err := appointmentvalidation.RunUpdateApptValidation(r, a)

		if err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprint(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), entity.CtxKey("appt"), appointment)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
