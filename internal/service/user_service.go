package service

import (
	"context"

	"github.com/sqoopdata/madoc/internal/domain/entity"
	"github.com/sqoopdata/madoc/internal/domain/repository/userstore"
)

type UserService struct {
	store userstore.UserStore
}

func NewUserService(s userstore.UserStore) *UserService {
	return &UserService{store: s}
}

func (s *UserService) AddUser(ctx context.Context, u *entity.User) error {
	return s.store.Add(ctx, u)
}

func (s *UserService) UpdateUser(ctx context.Context, u *entity.User) error {
	return s.store.Update(ctx, u)
}

func (s *UserService) GetUser(ctx context.Context, username string) (*entity.User, error) {
	return s.store.Get(ctx, username)
}

func (s *UserService) GetAllUsers(ctx context.Context) (*[]entity.User, error) {
	return s.store.GetAll(ctx)
}
