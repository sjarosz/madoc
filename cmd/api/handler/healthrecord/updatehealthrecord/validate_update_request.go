package updatehealthrecord

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/healthrecordvalidation"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func validateUpdateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		record, err := healthrecordvalidation.RunUpdateRecordValidation(r, a)

		if err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprint(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), entity.CtxKey("record"), record)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
