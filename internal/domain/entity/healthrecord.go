package entity

import (
	"time"
)

type HealthRecord struct {
	HealthRecordId int       `json:"healthRecordId"`
	ApptId         int       `json:"apptId"`
	Patient        string    `json:"patient"`
	Description    string    `json:"description"`
	CreatedBy      string    `json:"createdBy"`
	Created        time.Time `json:"created"`
}
