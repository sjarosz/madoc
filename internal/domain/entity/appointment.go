package entity

import (
	"time"
)

const (
	ACTIVE UType = iota + 1
	CANCELLED
	COMPLETED
)

type ApptStatus int

type Appointment struct {
	ApptId    int        `json:"apptId"`
	StartTime time.Time  `json:"startTime"`
	EndTime   time.Time  `json:"endTime"`
	Patient   string     `json:"patient"`
	Status    ApptStatus `json:"status"`
	CreatedBy string     `json:"createdBy"`
	Created   time.Time  `json:"created"`
}
