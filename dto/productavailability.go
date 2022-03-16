package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//ProductAvailability represents product availability details
type ProductAvailability struct {
	ProductCode                 string                       `json:"productCode"`
	ProductCategory             string                       `json:"productCategory"`
	LocationAvailabilityDetails []LocationAvailabilityDetail `json:"locationAvailabilityDetails,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ProductAvailability) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
