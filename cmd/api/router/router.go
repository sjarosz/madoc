package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/createappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/getappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/updateappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/createhealthrecord"
	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/gethealthrecord"
	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/updatehealthrecord"
	"github.com/sqoopdata/madoc/cmd/api/handler/index"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/createuser"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/getuser"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/updateuser"
	"github.com/sqoopdata/madoc/internal/application"
)

func Get(app *application.Application) *mux.Router {
	mux := mux.NewRouter()

	mux.Handle("/", index.HandleRequest(app)).Methods(http.MethodGet)
	mux.Handle("/users/{id}", getuser.HandleGetRequest(app)).Methods(http.MethodGet)
	mux.Handle("/users", getuser.HandleGetAllRequest(app)).Methods(http.MethodGet)
	mux.Handle("/users", createuser.HandleCreateRequest(app)).Methods(http.MethodPost)
	mux.Handle("/users/{id}", updateuser.HandleUpdateRequest(app)).Methods(http.MethodPut)

	mux.Handle("/appointments", createappt.HandleCreateRequest(app)).Methods(http.MethodPost)
	mux.Handle("/appointments", getappt.HandleGetAllRequest(app)).Queries("username", "{username}").Methods(http.MethodGet)
	mux.Handle("/appointments/{apptId}", getappt.HandleGetRequest(app)).Methods(http.MethodGet)
	mux.Handle("/appointments/{apptId}", updateappt.HandleUpdateRequest(app)).Methods(http.MethodPut)

	mux.Handle("/healthrecords", createhealthrecord.HandleCreateRequest(app)).Methods(http.MethodPost)
	mux.Handle("/healthrecords", gethealthrecord.HandleGetByPatientRequest(app)).Queries("username", "{username}").Methods(http.MethodGet)
	mux.Handle("/healthrecords/{hrId}", updatehealthrecord.HandleUpdateRequest(app)).Methods(http.MethodPut)

	mux.Handle("/metrics", promhttp.Handler())
	return mux
}
