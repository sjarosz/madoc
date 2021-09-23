package service

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/domain/repository/appointmentstore"
)

type AppointmentService struct {
	store appointmentstore.AppointmentStore
}

func NewAppointmentService(s appointmentstore.AppointmentStore) *AppointmentService {
	return &AppointmentService{store: s}
}

func (s *AppointmentService) AddAppointment(ctx context.Context, u *entity.Appointment) error {
	return s.store.Add(ctx, u)
}

func (s *AppointmentService) UpdateAppointment(ctx context.Context, u *entity.Appointment) (*entity.Appointment, error) {
	return s.store.Update(ctx, u)
}

func (s *AppointmentService) GetAppointment(ctx context.Context, apptId int) (*entity.Appointment, error) {
	return s.store.Get(ctx, apptId)
}

func (s *AppointmentService) GetAllAppointments(ctx context.Context, patient string) (*[]entity.Appointment, error) {
	return s.store.GetAll(ctx, patient)
}
