package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//ApplicableDiscounts represents details of applicable discounts
type ApplicableDiscounts struct {
	DiscountedPrice float32           `json:"discountedPrice,omitempty"`
	DiscountDetails []DiscountDetails `json:"discountDetails,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ApplicableDiscounts) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
