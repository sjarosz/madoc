package createuser

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/sqoopdata/madoc/cmd/api/handler/user/uservalidation"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

var IsAlphabets = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func validateCreateRequest(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := uservalidation.RunCreateUserValidation(r, a)

		if err != nil {
			w.WriteHeader((http.StatusPreconditionFailed))
			fmt.Fprint(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), entity.CtxKey("user"), user)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
