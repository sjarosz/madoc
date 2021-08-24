package updateuser

import (
	"context"
	"encoding/json"
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
		defer r.Body.Close()

		params := mux.Vars(r)
		username := params["id"]
		if !IsAlphabets(username) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed username")
			return
		}

		user := &model.User{}
		json.NewDecoder(r.Body).Decode(user)
		user.Username = username

		if !IsAlphabets(user.FName) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed first name")
			return
		}

		if !IsAlphabets(user.LName) {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprintf(w, "malformed last name")
			return
		}

		ctx := context.WithValue(r.Context(), model.CtxKey("user"), user)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
