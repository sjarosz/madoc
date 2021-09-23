package uservalidation

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/handler/common"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func FirstName(firstName string) error {
	if len(firstName) == 0 {
		return errors.New("first name cannot be empty")
	}

	if !common.IsAlphabets(firstName) {
		return errors.New("malformed first name")
	}

	return nil
}
func LastName(lastName string) error {
	if len(lastName) == 0 {
		return errors.New("last name cannot be empty")
	}

	if !common.IsAlphabets(lastName) {
		return errors.New("malformed last name")
	}

	return nil
}

func Username(username string) error {
	if len(username) == 0 {
		return errors.New("username cannot be empty")
	}

	if !common.IsAlphaNumeric(username) {
		return errors.New("malformed username")
	}

	if len(username) > 25 {
		return errors.New("username len cannot be greater than 25 chars")
	}

	return nil
}

func UType(uType int) error {
	if uType == 0 || uType > 3 {
		return errors.New("unknown user type")
	}
	return nil
}

func RunCreateUserValidation(r *http.Request, app *application.Application) (*entity.User, error) {
	defer r.Body.Close()
	user := &entity.User{}
	json.NewDecoder(r.Body).Decode(user)

	if err := Username(user.Username); err != nil {
		return nil, err
	}

	if err := UType(int(user.UserType)); err != nil {
		return nil, err
	}
	return user, nil
}

func RunUpdateUserValidation(r *http.Request, app *application.Application) (*entity.User, error) {
	defer r.Body.Close()
	user := &entity.User{}
	json.NewDecoder(r.Body).Decode(user)

	params := mux.Vars(r)
	username := params["id"]

	if err := Username(username); err != nil {
		return nil, err
	}

	user.Username = username

	if err := FirstName(user.FName); err != nil {
		return nil, err
	}

	if err := LastName(user.LName); err != nil {
		return nil, err
	}

	return user, nil
}
