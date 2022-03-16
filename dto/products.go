package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//Products represents list of products
type Products []Product

//FromBytes performs conversion of http response to the representing struct
func (dto *Products) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
