package model

import (
	"context"
	"time"

	"github.com/sqoopdata/madoc/pkg/application"
)

var (
	HEALTH_RECORD_INSERT_STMT        = "INSERT INTO healthRecords(apptId, description, patient, createdBy, created) VALUES($1, $2, $3, $4, current_timestamp) RETURNING healthRecordId, created"
	HEALTH_RECORD_SELECT_DISTINCT    = "SELECT DISTINCT healthRecordId, apptId, description, patient, createdBy, created FROM healthRecords WHERE healthRecordId = $1"
	HEALTH_RECORD_SELECT_BY_USERNAME = "SELECT healthRecordId, apptId, description, patient, createdBy, created FROM healthRecords WHERE patient = $1"
)

type HealthRecord struct {
	HealthRecordId int       `json:"healthRecordId"`
	ApptId         int       `json:"apptId"`
	Patient        string    `json:"patient"`
	Description    string    `json:"description"`
	CreatedBy      string    `json:"createdBy"`
	Created        time.Time `json:"created"`
}

func (r *HealthRecord) Create(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, HEALTH_RECORD_INSERT_STMT, &r.ApptId, &r.Description, &r.Patient, &r.CreatedBy).Scan(&r.HealthRecordId, &r.Created)

	if err != nil {
		return err
	}

	return nil
}

func (a *HealthRecord) GetByUsername(ctx context.Context, app *application.Application) ([]HealthRecord, error) {
	rows, err := app.DB.Client.QueryContext(ctx, HEALTH_RECORD_SELECT_BY_USERNAME, &a.Patient)

	if err != nil {
		return nil, err
	}

	var appts []HealthRecord

	defer rows.Close()
	for rows.Next() {
		var healthRecordId, apptId int
		var created time.Time
		var patient, description, createdBy string

		err := rows.Scan(&healthRecordId, &apptId, &description, &patient, &createdBy, &created)

		if err != nil {
			return nil, err
		}

		appts = append(appts, HealthRecord{HealthRecordId: healthRecordId, ApptId: apptId, Description: description, Patient: patient, CreatedBy: createdBy, Created: created})
	}

	return appts, nil
}

func (r *HealthRecord) GetByHealthRecordId(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, HEALTH_RECORD_SELECT_DISTINCT, &r.HealthRecordId).Scan(&r.HealthRecordId, &r.ApptId, &r.Description, &r.Patient, &r.CreatedBy, &r.Created)

	if err != nil {
		return err
	}

	return nil
}
