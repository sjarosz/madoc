package healthrecordstore

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

type HealthRecordStore interface {
	Add(ctx context.Context, u *entity.HealthRecord) (*entity.HealthRecord, error)
	Update(ctx context.Context, u *entity.HealthRecord) (*entity.HealthRecord, error)
	Get(ctx context.Context, patient string) (*[]entity.HealthRecord, error)
}
