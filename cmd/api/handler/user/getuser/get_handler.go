package getuser

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func getUser(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id := r.Context().Value(entity.CtxKey("id"))
		user, err := app.UserService.GetUser(r.Context(), id.(string))

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprint(w, "user does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Something went wrong. Try again!")
			return
		}

		response, _ := json.Marshal(user)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func HandleGetRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateGetRequest,
	}

	return middleware.Chain(getUser(app), app, mdw...)
}
