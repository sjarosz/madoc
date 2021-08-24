package createuser

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

		user := &model.User{}
		json.NewDecoder(r.Body).Decode(user)

		if !IsAlphabets(user.Username) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed username")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("user"), user)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
