package createhealthrecord

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

func validateCreateRules(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rO := r.Context().Value(model.CtxKey("record"))
		record := rO.(*model.HealthRecord)

		u := &model.User{Username: record.CreatedBy}
		u.GetUserByUsername(r.Context(), a)

		if u.UserType != model.DOCTOR {
			fmt.Fprintf(w, "you must be a doctor to create health records")
			return
		}

		appt := &model.Appointment{ApptId: record.ApptId}
		appt.GetByApptId(r.Context(), a)

		if record.Patient != appt.Patient {
			fmt.Fprintf(w, "appointment referenced is for another patient '%s'. you cannot hijack someone else's appt.", appt.Patient)
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("record"), record)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
