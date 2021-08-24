package getuser

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/cmd/api/model"
	"github.com/sqoopdata/madoc/pkg/application"
	"github.com/sqoopdata/madoc/pkg/middleware"
)

func getUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id := r.Context().Value(model.CtxKey("id"))
		user := &model.User{Username: id.(string)}

		if err := user.GetUserByUsername(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "user does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		response, _ := json.Marshal(user)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func HandleRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateRequest,
	}

	return middleware.Chain(getUser(app), app, mdw...)
}
