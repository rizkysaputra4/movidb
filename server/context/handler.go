package context

import (
	"encoding/json"
	"net/http"
)

// Context ...
type Context struct {
	Res  http.ResponseWriter
	Req  *http.Request
	Data interface{}
}

// JSONDecoder is decoding request data into structure
func (c Context) JSONDecoder() error {
	if err := json.NewDecoder(c.Req.Body).Decode(c.Data); err != nil {
		c.SendError(http.StatusInternalServerError, err.Error(), "Error when decoding the request body")
		return err
	}
	return nil
}

// ErrorUpdatingDB ...
func (c Context) ErrorUpdatingDB(err error) {
	c.SendError(http.StatusInternalServerError, err.Error(), "Error when updating data into db")
}

// ErrorInsertingDataIntoDB ...
func (c Context) ErrorInsertingDataIntoDB(err error) {
	c.SendError(http.StatusInternalServerError, err.Error(), "Error when inserting data into db")
}

// ErrorDeletingDataFromDB ...
func (c Context) ErrorDeletingDataFromDB(err error) {
	c.SendError(http.StatusInternalServerError, err.Error(), "Error when deleting data from db")
}

// ErrorGettingDataFromDB ...
func (c Context) ErrorGettingDataFromDB(err error) {
	c.SendError(http.StatusInternalServerError, err.Error(), "Error when getting data from db")
}

// SendError ...
func (c Context) SendError(code int, err string, message string) {
	data := &StatusResponse{
		Status:  code,
		Error:   err,
		Message: message,
	}

	response, _ := json.Marshal(data)
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(http.StatusOK)
	c.Res.Write(response)
}

// SendSuccess ...
func (c Context) SendSuccess(message ...interface{}) {
	data := &StatusResponse{
		Status: http.StatusOK,
		Error:  "OK",
		Data:   c.Data,
	}

	if len(message) > 1 {
		data.Data = message
	} else if len(message) == 1 {
		data.Data = message[0]
	}

	response, _ := json.Marshal(data)
	c.Res.Header().Set("Content-Type", "application/json")
	c.Res.WriteHeader(http.StatusOK)
	c.Res.Write(response)

}

// StatusResponse ...
type StatusResponse struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
