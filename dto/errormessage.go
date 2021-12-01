package dto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ErrorMessage represents API response error messages.
type ErrorMessage struct {
	Message          string   `json:"message"`
	ValidationErrors []string `json:"validationErrors"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ErrorMessage) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	} 
	return err
}
