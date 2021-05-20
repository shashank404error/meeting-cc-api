package data

import (
	"github.com/go-playground/validator/v10"
	
	//"github.com/bybrisk/structs"
)

//request to create event
type CreateEventRequest struct {
	//Summery of the event
	//
	//required: true
	Summary string `json:"summary" validate:"required"`

	//Location of the event
	//
	//required: false
	Location string `json:"location"`

	//Description of the event
	//
	//required: true
	Description string `json:"description" validate:"required"`

	//user primary emailID
	//
	//required: true
	Email string `json:"email" validate:"required"`

	//Start time in RFC3339 format
	//
	//required: true
	StartTime  string `json:"startTime" validate:"required"`

	//End time in RFC3339 format
	//
	//required: true
	EndTime  string `json:"endTime" validate:"required"`

	//timezone of the user
	//
	//required: false
	TimeZone string `json:"timezone"`
}

// response of create event
type CreateEventResponse struct {

	//Resonse Code
	//
	Code int64 `json:"code"`
}

//Response of a event details
type GetEventDetailResponse struct {

	//Resonse Code
	//
	Code int64 `json:"code"`
}

func (d *CreateEventRequest) ValidateCreateEventRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}