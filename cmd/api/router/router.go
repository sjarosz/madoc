package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/createappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/deleteappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/appointment/getappt"
	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/createhealthrecord"
	"github.com/sqoopdata/madoc/cmd/api/handler/healthrecord/gethealthrecord"
	"github.com/sqoopdata/madoc/cmd/api/handler/index"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/createuser"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/getuser"
	"github.com/sqoopdata/madoc/cmd/api/handler/user/updateuser"
	"github.com/sqoopdata/madoc/pkg/application"
)

func Get(app *application.Application) *mux.Router {
	mux := mux.NewRouter()

	mux.Handle("/", index.HandleRequest(app)).Methods(http.MethodGet)
	mux.Handle("/users/{id}", getuser.HandleRequest(app)).Methods(http.MethodGet)
	mux.Handle("/users", getuser.HandleRequestAll(app)).Methods(http.MethodGet)
	mux.Handle("/users", createuser.HandleRequest(app)).Methods(http.MethodPost)
	mux.Handle("/users/{id}", updateuser.HandleRequest(app)).Methods(http.MethodPut)

	mux.Handle("/appointments", createappt.HandleRequest(app)).Methods(http.MethodPost)
	mux.Handle("/appointments", getappt.HandleRequestByUsername(app)).Queries("username", "{username}").Methods(http.MethodGet)
	mux.Handle("/appointments/{apptId}", getappt.HandleRequestByApptId(app)).Methods(http.MethodGet)
	mux.Handle("/appointments", deleteappt.HandleRequest(app)).Queries("apptId", "{apptId}").Methods(http.MethodDelete)

	mux.Handle("/healthrecords", createhealthrecord.HandleRequest(app)).Methods(http.MethodPost)
	mux.Handle("/healthrecords", gethealthrecord.HandleRequestByUsername(app)).Queries("username", "{username}").Methods(http.MethodGet)
	mux.Handle("/healthrecords/{healthRecordId}", gethealthrecord.HandleRequestByHealthRecordId(app)).Methods(http.MethodGet)

	mux.Handle("/metrics", promhttp.Handler())
	return mux
}
