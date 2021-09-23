package service

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/domain/repository/healthrecordstore"
)

type HealthRecordService struct {
	store healthrecordstore.HealthRecordStore
}

func NewHealthRecordService(s healthrecordstore.HealthRecordStore) *HealthRecordService {
	return &HealthRecordService{store: s}
}

func (s *HealthRecordService) AddHealthRecord(ctx context.Context, u *entity.HealthRecord) (*entity.HealthRecord, error) {
	return s.store.Add(ctx, u)
}

func (s *HealthRecordService) UpdateHealthRecord(ctx context.Context, u *entity.HealthRecord) (*entity.HealthRecord, error) {
	return s.store.Update(ctx, u)
}

func (s *HealthRecordService) GetByPatient(ctx context.Context, patient string) (*[]entity.HealthRecord, error) {
	return s.store.Get(ctx, patient)
}
