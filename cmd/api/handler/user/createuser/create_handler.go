package createuser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/middleware"
)

func create(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		uObj := r.Context().Value(entity.CtxKey("user"))
		user := uObj.(*entity.User)

		if err := app.UserService.AddUser(r.Context(), user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		response, _ := json.Marshal(user)
		w.Write(response)
	}
}

func HandleCreateRequest(app *application.Application) http.HandlerFunc {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		middleware.SecureHeaders,
		validateCreateRequest,
	}

	return middleware.Chain(create(app), app, mdw...)
}
