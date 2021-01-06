package model

// StatusResponse ...
type StatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	/* Error code details
	0: success
	1:
	*/
}
