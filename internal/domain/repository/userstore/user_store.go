package userstore

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

type UserStore interface {
	Add(ctx context.Context, u *entity.User) error
	Update(ctx context.Context, u *entity.User) error
	Get(ctx context.Context, username string) (*entity.User, error)
	GetAll(ctx context.Context) (*[]entity.User, error)
}
