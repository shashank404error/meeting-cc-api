package data

import (
	//"github.com/go-playground/validator/v10"
	//"github.com/bybrisk/structs"
)

//Response of a event details
type GetEventDetailResponse struct {

	//Resonse Code
	// Example: 200 for success
	Code int64 `json:"code"`
}

/*func (d *UpdateDeliveryDistance) ValidateUpdateDeliveryDistance() error {
	validate := validator.New()
	return validate.Struct(d)
}*/