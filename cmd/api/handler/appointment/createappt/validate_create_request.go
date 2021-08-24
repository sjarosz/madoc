package createappt

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func validateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		appointment := &model.Appointment{}
		json.NewDecoder(r.Body).Decode(appointment)

		if time.Now().After(appointment.StartTime) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "appointment start time must be in the future")
			return
		}

		if appointment.StartTime.After(appointment.EndTime) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "appointment start time must be before end time")
			return
		}

		if appointment.EndTime.Sub(appointment.StartTime).Minutes() > 30 {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "appointment cannot exceed 30 mins window")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("appt"), appointment)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
