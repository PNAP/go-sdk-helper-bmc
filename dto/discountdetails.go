package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//DiscountDetails represents details of the discount
type DiscountDetails struct {
	Code  string  `json:"code"`
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *DiscountDetails) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
