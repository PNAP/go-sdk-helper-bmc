package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

// ApplicableDiscountDetails represents details of the discount
type ApplicableDiscountDetails struct {
	Code       string  `json:"code"`
	Type       string  `json:"type"`
	Value      float32 `json:"value"`
	CouponCode *string `json:"couponCode,omitempty"`
}

// FromBytes performs conversion of http response to the representing struct
func (dto *ApplicableDiscountDetails) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
