package userstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/sqoopdata/madoc/internal/domain/entity"
)

var (
	USER_INSERT_STMT     = "INSERT INTO users(username, utype, fName, lName) VALUES($1, $2, $3, $4) RETURNING id, username, created"
	USER_UPDATE_STMT     = "UPDATE users SET fName = $2, lName = $3 WHERE username = $1 RETURNING id, username, utype, fName, lName, created"
	USER_SELECT_DISTINCT = "SELECT DISTINCT id, username, utype, fName, lName, created FROM users WHERE username = $1"
	USER_SELECT_ALL      = "SELECT id, username, utype, fName, lName, created FROM users"
)

type PGUserStore struct {
	store *sql.DB
}

// Returns concrete (struct) type so that consumer can directly call functions
func NewUserStore(db *sql.DB) *PGUserStore {
	return &PGUserStore{store: db}
}

func (s *PGUserStore) Add(ctx context.Context, u *entity.User) error {
	err := s.store.QueryRowContext(ctx, USER_INSERT_STMT, u.Username, u.UserType, u.FName, u.LName).Scan(&u.Id, &u.Username, &u.Created)

	if err != nil {
		return err
	}

	return nil
}

func (s *PGUserStore) Update(ctx context.Context, u *entity.User) error {
	err := s.store.QueryRowContext(ctx, USER_UPDATE_STMT, u.Username, u.FName, u.LName).Scan(&u.Id, &u.Username, &u.UserType, &u.FName, &u.LName, &u.Created)

	if err != nil {
		return err
	}

	return nil
}

func (s *PGUserStore) Get(ctx context.Context, username string) (*entity.User, error) {
	var u entity.User
	err := s.store.QueryRowContext(ctx, USER_SELECT_DISTINCT, username).Scan(&u.Id, &u.Username, &u.UserType, &u.FName, &u.LName, &u.Created)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *PGUserStore) GetAll(ctx context.Context) (*[]entity.User, error) {
	rows, err := s.store.QueryContext(ctx, USER_SELECT_ALL)

	if err != nil {
		return nil, err
	}

	var users []entity.User

	defer rows.Close()
	for rows.Next() {
		var id int
		var username, fName, lName string
		var created time.Time
		var uType entity.UType

		err := rows.Scan(&id, &username, &uType, &fName, &lName, &created)

		if err != nil {
			return nil, err
		}

		users = append(users, entity.User{Id: id, Username: username, UserType: uType, FName: fName, LName: lName, Created: created})
	}

	return &users, nil
}
