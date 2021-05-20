package data

import (
	"encoding/json"
	"io"
)	

func (d *GetEventDetailResponse) GetEventDetailResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

func (d *CreateEventRequest) FromJSONToCreateEventRequest (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}

func (d *CreateEventResponse) CreateEventResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}