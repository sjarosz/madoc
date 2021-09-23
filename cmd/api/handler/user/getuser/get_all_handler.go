package getuser

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func getAllUsers(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		users, err := app.UserService.GetAllUsers(r.Context())

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprint(w, "no users found")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Something went wrong. Try again!")
			return
		}

		response, _ := json.Marshal(users)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func HandleGetAllRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
	}

	return middleware.Chain(getAllUsers(app), app, mdw...)
}
