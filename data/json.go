package data

import (
	"encoding/json"
	"io"
)	

func (d *GetEventDetailResponse) GetEventDetailResponseToJSON (w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(d)
}

/*func (d *AddDeliveryRequestWithGeoCode) FromJSONToAddDeliveryStruct (r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(d)
}*/