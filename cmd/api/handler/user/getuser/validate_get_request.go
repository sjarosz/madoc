package getuser

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/uservalidation"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func validateGetRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		if err := uservalidation.Username(id); err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprint(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), entity.CtxKey("id"), id)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
