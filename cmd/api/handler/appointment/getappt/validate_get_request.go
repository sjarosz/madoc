package getappt

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func validateGetRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		apptId := params["apptId"]
		id, err := strconv.Atoi(apptId)

		if err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprint(w, "malformed apptId")
			return

		}

		ctx := context.WithValue(r.Context(), entity.CtxKey("apptId"), id)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
