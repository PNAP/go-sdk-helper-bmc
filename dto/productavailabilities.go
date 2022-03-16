package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//ProductAvailabilities represents list of product availabilities
type ProductAvailabilities []ProductAvailability

//FromBytes performs conversion of http response to the representing struct
func (dto *ProductAvailabilities) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
