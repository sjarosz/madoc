package model

import (
	"context"
	"time"

	"github.com/sqoopdata/madoc/pkg/application"
)

var (
	USER_INSERT_STMT     = "INSERT INTO users(username, utype, fName, lName) VALUES($1, $2, $3, $4) RETURNING id, username, created"
	USER_UPDATE_STMT     = "UPDATE users SET fName = $2, lName = $3 WHERE username = $1 RETURNING id, username, utype, fName, lName, created"
	USER_SELECT_DISTINCT = "SELECT DISTINCT id, username, utype, fName, lName, created FROM users WHERE username = $1"
	USER_SELECT_ALL      = "SELECT id, username, utype, fName, lName, created FROM users"
)

type UType int

const (
	ADMIN UType = iota + 1
	PATIENT
	DOCTOR
)

// User captures identity information of a given user
type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	FName    string    `json:"firstName"`
	LName    string    `json:"lastName"`
	Created  time.Time `json:"created"`
	UserType UType     `json:"utype"`
}

// Create creates user in the db
func (u *User) Create(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, USER_INSERT_STMT, u.Username, u.UserType, u.FName, u.LName).Scan(&u.Id, &u.Username, &u.Created)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByUsername(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, USER_SELECT_DISTINCT, u.Username).Scan(&u.Id, &u.Username, &u.UserType, &u.FName, &u.LName, &u.Created)

	if err != nil {
		return err
	}

	return nil
}
func GetAllUsers(ctx context.Context, app *application.Application) ([]User, error) {
	rows, err := app.DB.Client.QueryContext(ctx, USER_SELECT_ALL)

	if err != nil {
		return nil, err
	}

	var users []User

	defer rows.Close()
	for rows.Next() {
		var id int
		var username, fName, lName string
		var created time.Time
		var uType UType

		err := rows.Scan(&id, &username, &uType, &fName, &lName, &created)

		if err != nil {
			return nil, err
		}

		users = append(users, User{Id: id, Username: username, UserType: uType, FName: fName, LName: lName, Created: created})
	}

	return users, nil
}
func (u *User) Update(ctx context.Context, app *application.Application) error {
	err := app.DB.Client.QueryRowContext(ctx, USER_UPDATE_STMT, u.Username, u.FName, u.LName).Scan(&u.Id, &u.Username, &u.UserType, &u.FName, &u.LName, &u.Created)

	if err != nil {
		return err
	}

	return nil
}
