package getuser

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
)

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func validateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		if !IsAlphabets(id) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("id"), id)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
