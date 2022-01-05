package dto

import (
	"fmt"
	"net/http"
)

//ErrorResolver represents instance that handles processing of errors and creates error messages.
type ErrorResolver struct {
	Error error
}

//NewErrorResolver constructs new instance of ErrorProcessing struct
func NewErrorResolver(httpResponse *http.Response, err error) *ErrorResolver {

	if err != nil && httpResponse == nil {
		//do nothing  use existing err
	} else if err != nil {
		response := &ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			//do nothing use existing err
		} else {
			err = fmt.Errorf("returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
		}
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		//should not happen ever
		err = nil
	} else {
		response := &ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			err = error
		} else {
			err = fmt.Errorf("returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
		}

	}
	return &ErrorResolver{Error: err}
}
