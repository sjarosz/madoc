package appointmentstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

var (
	APPT_INSERT_STMT     = "INSERT INTO appointments(startTime, endTime, patient, status, createdBy, created) VALUES($1, $2, $3, 1, $4, current_timestamp) RETURNING apptId, startTime, endTime, patient, status, createdBy, created"
	APPT_SELECT_DISTINCT = "SELECT DISTINCT apptId, startTime, endTime, patient, status, createdBy, created FROM appointments WHERE apptId = $1"
	APPT_SELECT_ALL      = "SELECT * FROM APPOINTMENTS WHERE patient = $1 and startTime >= now()"
	APPT_UPDATE          = "UPDATE appointments SET startTime = $2, endTime = $3, status = $4 WHERE apptId = $1 RETURNING apptId, startTime, endTime, patient, status, createdBy, created"
)

type PGAppointmentStore struct {
	store *sql.DB
}

// Returns concrete (struct) type so that consumer can directly call functions
func NewAppointmentStore(db *sql.DB) *PGAppointmentStore {
	return &PGAppointmentStore{store: db}
}

func (s *PGAppointmentStore) Add(ctx context.Context, a *entity.Appointment) error {
	err := s.store.QueryRowContext(ctx, APPT_INSERT_STMT, &a.StartTime, &a.EndTime, &a.Patient, &a.CreatedBy).Scan(&a.ApptId, &a.StartTime, &a.EndTime, &a.Patient, &a.Status, &a.CreatedBy, &a.Created)

	if err != nil {
		return err
	}

	return nil
}

func (s *PGAppointmentStore) Get(ctx context.Context, id int) (*entity.Appointment, error) {
	var a entity.Appointment
	err := s.store.QueryRowContext(ctx, APPT_SELECT_DISTINCT, id).Scan(&a.ApptId, &a.StartTime, &a.EndTime, &a.Patient, &a.Status, &a.CreatedBy, &a.Created)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
func (s *PGAppointmentStore) Update(ctx context.Context, a *entity.Appointment) (*entity.Appointment, error) {
	var cAppt entity.Appointment
	err := s.store.QueryRowContext(ctx, APPT_UPDATE, a.ApptId, a.StartTime, a.EndTime, a.Status).Scan(&cAppt.ApptId, &cAppt.StartTime, &cAppt.EndTime, &cAppt.Patient, &cAppt.Status, &cAppt.CreatedBy, &cAppt.Created)

	if err != nil {
		return nil, err
	}

	return &cAppt, nil
}

func (s *PGAppointmentStore) GetAll(ctx context.Context, patient string) (*[]entity.Appointment, error) {
	rows, err := s.store.QueryContext(ctx, APPT_SELECT_ALL, patient)

	if err != nil {
		return nil, err
	}

	var appointments []entity.Appointment

	defer rows.Close()
	for rows.Next() {
		var apptId int
		var status entity.ApptStatus
		var startTime, endTime, created time.Time
		var createdBy, patient string

		err := rows.Scan(&apptId, &startTime, &endTime, &patient, &status, &createdBy, &created)

		if err != nil {
			return nil, err
		}

		appointments = append(appointments, entity.Appointment{ApptId: apptId, StartTime: startTime, EndTime: endTime, Patient: patient, CreatedBy: createdBy, Created: created, Status: status})
	}

	return &appointments, nil
}
