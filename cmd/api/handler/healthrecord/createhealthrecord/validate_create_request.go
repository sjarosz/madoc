package createhealthrecord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func validateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		record := &model.HealthRecord{}
		json.NewDecoder(r.Body).Decode(record)

		if len(record.Description) < 1 {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "record description cannot be empty")
			return
		}

		if len(record.Patient) < 1 {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "record patient cannot be empty")
			return
		}
		ctx := context.WithValue(r.Context(), model.CtxKey("record"), record)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
