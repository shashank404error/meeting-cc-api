package handlers

import (
	"net/http"
	"fmt"
	"github.com/shashank404error/meeting-cc-api/data"
)

// swagger:route POST /meeting/create/callback meeting createEvent
// Create events in the primary calender of the user
//
// responses:
//	200: createEventResp
//  422: errorValidation
//  501: errorResponse

func (p *Meeting) CreateEvent (w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request -> meeting-cc-api Module")
	event := &data.CreateEventRequest{}

	err:=event.FromJSONToCreateEventRequest(r.Body)
	if err!=nil {
		http.Error(w,"Data failed to unmarshel", http.StatusBadRequest)
	}

	//validate the data
	err = event.ValidateCreateEventRequest()
	if err!=nil {
		p.l.Println("Validation error in POST request -> meeting-cc-api Module \n",err)
		http.Error(w,fmt.Sprintf("Error in data validation : %s",err), http.StatusBadRequest)
		return
	} 

	//add request to event API
	res := data.CreateEventCRUDOPS(event,r)

	//writing to the io.Writer
	err = res.CreateEventResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data with ID failed to marshel",http.StatusInternalServerError)		
	}
}