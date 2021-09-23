package healthrecordvalidation

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/cmd/api/handler/common"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func Description(description string) error {
	if len(description) == 0 {
		return errors.New("description cannot be empty")
	}

	return nil
}

func Patient(patient string) error {
	if len(patient) == 0 {
		return errors.New("patient cannot be empty")
	}

	if !common.IsAlphabets(patient) {
		return errors.New("malformed patient")
	}

	return nil
}

func Id(id int) error {
	if id == 0 {
		return errors.New("health record id is not valid")
	}
	return nil
}

func CreatedBy(createdBy string) error {
	if len(createdBy) == 0 {
		return errors.New("createdBy is not valid")
	}

	if !common.IsAlphabets(createdBy) {
		return errors.New("malformed createdBy")
	}

	return nil
}

func UserMustBeDoctor(record *entity.HealthRecord, c context.Context, app *application.Application) error {
	u, err := app.UserService.GetUser(c, record.CreatedBy)

	if err != nil {
		return err
	}

	if u.UserType != entity.DOCTOR {
		return errors.New("you must be doctor to create health records")
	}

	return nil
}

func UserMustUseSelfAppt(record *entity.HealthRecord, c context.Context, app *application.Application) error {
	a, err := app.AppointmentService.GetAppointment(c, record.ApptId)

	if err != nil {
		return err
	}

	if a.Patient != record.Patient {
		return errors.New("cannot use someone else appointment")
	}

	return nil
}

func RunCreateRecordValidation(r *http.Request, app *application.Application) (*entity.HealthRecord, error) {
	defer r.Body.Close()

	record := &entity.HealthRecord{}
	json.NewDecoder(r.Body).Decode(record)

	if err := Patient(record.Patient); err != nil {
		return nil, err
	}

	if err := CreatedBy(record.CreatedBy); err != nil {
		return nil, err
	}

	if err := Description(record.Description); err != nil {
		return nil, err
	}

	if err := UserMustBeDoctor(record, r.Context(), app); err != nil {
		return nil, err
	}

	if err := UserMustUseSelfAppt(record, r.Context(), app); err != nil {
		return nil, err
	}

	return record, nil
}

func RunUpdateRecordValidation(r *http.Request, app *application.Application) (*entity.HealthRecord, error) {
	defer r.Body.Close()
	record := &entity.HealthRecord{}
	json.NewDecoder(r.Body).Decode(record)

	params := mux.Vars(r)
	hrId, err := strconv.Atoi(params["hrId"])

	if err != nil {
		return nil, err
	}

	if err := Id(hrId); err != nil {
		return nil, err
	}

	record.HealthRecordId = hrId

	if err := Patient(record.Patient); err != nil {
		return nil, err
	}

	if err := Description(record.Description); err != nil {
		return nil, err
	}

	return record, nil
}
