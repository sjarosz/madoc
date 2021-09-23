package appointmentvalidation

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sqoopdata/madoc/internal/application"
	"github.com/sqoopdata/madoc/internal/domain/entity"
)

func Time(start time.Time, end time.Time) error {
	if time.Now().After(start) {
		return errors.New("appointment start time must be in the future")
	}
	if start.After(end) {
		return errors.New("appointment start time must be before end time")
	}
	if end.Before(start) {
		return errors.New("appointment end time must be after start time")
	}
	if end.Sub(start).Minutes() > 30 {
		return errors.New("appointment cannot exceed 30 mins window")
	}

	return nil
}

func ApptId(apptId int) error {
	if apptId == 0 {
		return errors.New("malformed appointment id")
	}

	return nil
}

func RunUpdateApptValidation(r *http.Request, app *application.Application) (*entity.Appointment, error) {
	defer r.Body.Close()

	appointment := &entity.Appointment{}
	json.NewDecoder(r.Body).Decode(appointment)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["apptId"])

	if err != nil {
		return nil, err
	}

	if err := ApptId(id); err != nil {
		return nil, err
	}

	if err := Time(appointment.StartTime, appointment.EndTime); err != nil {
		return nil, err
	}

	appointment.ApptId = id

	return appointment, nil
}

func RunCreateApptValidation(r *http.Request, app *application.Application) (*entity.Appointment, error) {
	defer r.Body.Close()

	appointment := &entity.Appointment{}
	json.NewDecoder(r.Body).Decode(appointment)

	if err := Time(appointment.StartTime, appointment.EndTime); err != nil {
		return nil, err
	}

	return appointment, nil
}
