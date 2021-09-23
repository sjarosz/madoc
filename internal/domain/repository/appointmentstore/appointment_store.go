package appointmentstore

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

type AppointmentStore interface {
	Add(ctx context.Context, a *entity.Appointment) error
	Update(ctx context.Context, a *entity.Appointment) (*entity.Appointment, error)
	Get(ctx context.Context, id int) (*entity.Appointment, error)
	GetAll(ctx context.Context, patient string) (*[]entity.Appointment, error)
}
