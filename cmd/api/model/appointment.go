package model

import (
	"context"
	"errors"
	"time"

	"github.com/sqoopdata/madoc/pkg/application"
)

var (
	APPT_INSERT_STMT        = "INSERT INTO appointments(startTime, endTime, patient, createdBy, created) VALUES($1, $2, $3, $4, current_timestamp) RETURNING apptId, created"
	APPT_SELECT_DISTINCT    = "SELECT DISTINCT apptId, startTime, endTime, patient, createdBy, created FROM appointments WHERE apptId = $1"
	APPT_SELECT_BY_USERNAME = "SELECT apptId, startTime, endTime, patient, createdBy, created FROM appointments WHERE patient = $1"
	APPT_DELETE             = "DELETE FROM APPOINTMENTS WHERE apptId = $1"
)

type Appointment struct {
	ApptId    int       `json:"apptId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Patient   string    `json:"patient"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
}

func (a *Appointment) Create(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, APPT_INSERT_STMT, &a.StartTime, &a.EndTime, &a.Patient, &a.CreatedBy).Scan(&a.ApptId, &a.Created)

	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) GetByUsername(ctx context.Context, app *application.Application) ([]Appointment, error) {
	rows, err := app.DB.Client.QueryContext(ctx, APPT_SELECT_BY_USERNAME, &a.Patient)

	if err != nil {
		return nil, err
	}

	var appts []Appointment

	defer rows.Close()
	for rows.Next() {
		var apptId int
		var startTime, endTime, created time.Time
		var patient, createdBy string

		err := rows.Scan(&apptId, &startTime, &endTime, &patient, &createdBy, &created)

		if err != nil {
			return nil, err
		}

		appts = append(appts, Appointment{ApptId: apptId, StartTime: startTime, EndTime: endTime, Patient: patient, CreatedBy: createdBy, Created: created})
	}

	return appts, nil
}

func (a *Appointment) GetByApptId(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, APPT_SELECT_DISTINCT, &a.ApptId).Scan(&a.ApptId, &a.StartTime, &a.EndTime, &a.Patient, &a.CreatedBy, &a.Created)

	if err != nil {
		return err
	}

	return nil
}

func (a *Appointment) DeleteByApptId(ctx context.Context, app *application.Application) error {
	result, err := app.DB.Client.ExecContext(ctx, APPT_DELETE, &a.ApptId)

	if err != nil {
		return err
	}

	n, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if n < 1 {
		return errors.New("Delete operation may have failed")
	}

	return nil
}
