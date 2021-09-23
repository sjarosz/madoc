package application

import (
	"database/sql"

	"github.com/sqoopdata/madoc/internal/config"
	"github.com/sqoopdata/madoc/internal/db"
	"github.com/sqoopdata/madoc/internal/domain/repository/appointmentstore"
	"github.com/sqoopdata/madoc/internal/domain/repository/healthrecordstore"
	"github.com/sqoopdata/madoc/internal/domain/repository/userstore"
	"github.com/sqoopdata/madoc/internal/service"
)

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB                  *db.DB
	Cfg                 *config.Config
	UserService         *service.UserService
	AppointmentService  *service.AppointmentService
	HealthRecordService *service.HealthRecordService
}

// Get captures env vars, establishes DB connection and keeps/returns
// reference to both
func Get() (*Application, error) {
	cfg := config.Get()

	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &Application{
		DB:                  db,
		Cfg:                 cfg,
		UserService:         initUserService(db.Client),
		HealthRecordService: initHealthRecordService(db.Client),
		AppointmentService:  initAppointmentService(db.Client),
	}, nil
}

func initUserService(client *sql.DB) *service.UserService {
	return service.NewUserService(userstore.NewUserStore(client))
}
func initAppointmentService(client *sql.DB) *service.AppointmentService {
	return service.NewAppointmentService(appointmentstore.NewAppointmentStore(client))
}
func initHealthRecordService(client *sql.DB) *service.HealthRecordService {
	return service.NewHealthRecordService(healthrecordstore.NewHealthRecordStore(client))
}
