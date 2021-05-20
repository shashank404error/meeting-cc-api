package data

import "net/http"

func GetEventsCRUDOPS (r *http.Request) *GetEventDetailResponse{
	EventDetailCallback(r)
	//sending response
	var response = GetEventDetailResponse{
		Code: 200,
	}

	return &response
}
