// Package classification of Meeting API
//
// Documentation for Meeting API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/shashank404error/meeting-cc-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters OAuth2
type callLoginParamsWrapper struct {
	// No data is required for loggin in
	// Call the API directly
}

// swagger:parameters getAllEvents
type callbackLoginParamsWrapper struct {
	// Callback of the meeting login API
	// gets called automatically
}

// Response structure for event detail
// swagger:response eventDetails
type getAllEventsResponseWrapper struct {
	// Response structre for all events detail
	// in: body
	Body data.GetEventDetailResponse
}

// swagger:parameters createEvent
type createEventParamsWrapper struct {
	// Structure to create event in primary calendar of the user
	//
	// in: body
	// required: true
	Body data.CreateEventRequest
}

// Response structure for create event 
// swagger:response createEventResp
type createEventsResponseWrapper struct {
	// Response structre for create event
	// in: body
	Body data.CreateEventResponse
}