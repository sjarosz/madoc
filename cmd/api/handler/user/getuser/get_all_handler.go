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

func getAllUsers(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		users, err := model.GetAllUsers(r.Context(), app)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "no users found")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, err.Error())
			return
		}

		response, _ := json.Marshal(users)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func HandleRequestAll(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
	}

	return middleware.Chain(getAllUsers(app), app, mdw...)
}
