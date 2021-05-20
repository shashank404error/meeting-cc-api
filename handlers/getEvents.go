package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shashank404error/meeting-cc-api/data"
)

// swagger:route GET /meeting/get/all/callback meeting getAllEvents
// Callback API for getting all events
//
// responses:
//	200: eventDetails
//  422: errorValidation
//  501: errorResponse

func (p *Delivery) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET request -> meeting-cc-api Module")

	lp := data.GetEventsCRUDOPS(r)

	err := lp.GetEventDetailResponseToJSON(w)
	if err!=nil {
		http.Error(w,"Data failed to marshel",http.StatusInternalServerError)		
	}
}