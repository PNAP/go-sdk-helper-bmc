package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//Product represents product details
type Product struct {
	ProductCode     string                 `json:"productCode"`
	ProductCategory string                 `json:"productCategory"`
	Plans           []PricingPlan          `json:"plans,omitempty"`
	Metadata        *ServerProductMetadata `json:"metadata,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *Product) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
