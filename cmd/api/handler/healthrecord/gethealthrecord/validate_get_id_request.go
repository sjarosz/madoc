package gethealthrecord

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

var IsNumeric = regexp.MustCompile(`^[0-9]+$`).MatchString

func validateRequestById(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		recordId := params["healthRecordId"]

		id, err := strconv.Atoi(recordId)
		if err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed recordId")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("recordId"), id)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
