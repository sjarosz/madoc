package healthrecordstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

var (
	HEALTH_RECORD_INSERT_STMT = "INSERT INTO healthrecords(apptId, description, patient, createdBy, created) VALUES($1, $2, $3, $4, current_timestamp) RETURNING healthRecordId, apptId, description, patient, createdBy, created"
	HEALTH_RECORD_BY_PATIENT  = "SELECT * FROM healthrecords WHERE patient = $1"
	HEALTH_RECORD_UPDATE      = "UPDATE healthrecords SET description = $2 WHERE healthRecordId = $1"
)

type PGHealthRecordStore struct {
	store *sql.DB
}

// Returns concrete (struct) type so that consumer can directly call functions
func NewHealthRecordStore(db *sql.DB) *PGHealthRecordStore {
	return &PGHealthRecordStore{store: db}
}

func (s *PGHealthRecordStore) Add(ctx context.Context, h *entity.HealthRecord) (*entity.HealthRecord, error) {
	var hr entity.HealthRecord
	err := s.store.QueryRowContext(ctx, HEALTH_RECORD_INSERT_STMT, &h.ApptId, &h.Description, &h.Patient, &h.CreatedBy).Scan(&hr.HealthRecordId, &hr.ApptId, &hr.Description, &hr.Patient, &hr.CreatedBy, &hr.Created)

	if err != nil {
		return nil, err
	}

	return &hr, nil
}

func (s *PGHealthRecordStore) Update(ctx context.Context, h *entity.HealthRecord) (*entity.HealthRecord, error) {
	_, err := s.store.QueryContext(ctx, HEALTH_RECORD_UPDATE, &h.ApptId, &h.Description)

	if err != nil {
		return nil, err
	}

	return h, nil
}

func (s *PGHealthRecordStore) Get(ctx context.Context, patient string) (*[]entity.HealthRecord, error) {
	rows, err := s.store.QueryContext(ctx, HEALTH_RECORD_BY_PATIENT, patient)

	if err != nil {
		return nil, err
	}

	var hrs []entity.HealthRecord

	defer rows.Close()
	for rows.Next() {
		var hId, apptId int
		var created time.Time
		var patient, description, createdBy string

		err := rows.Scan(&hId, &apptId, &description, &patient, &createdBy, &created)

		if err != nil {
			return nil, err
		}

		hrs = append(hrs, entity.HealthRecord{HealthRecordId: hId, ApptId: apptId, Description: description, Patient: patient, CreatedBy: createdBy, Created: created})
	}

	return &hrs, nil
}
