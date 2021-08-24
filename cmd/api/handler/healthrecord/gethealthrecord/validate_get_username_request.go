package gethealthrecord

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]{1,20}$`).MatchString

func validateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		username := params["username"]
		if !IsAlphabets(username) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed username")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("username"), username)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
