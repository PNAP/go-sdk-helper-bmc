package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//LocationAvailabilityDetail represents info about location, solutions and availability for a product
type LocationAvailabilityDetail struct {
	Location             string   `json:"location"`
	MinQuantityRequested int32    `json:"minQuantityRequested"`
	MinQuantityAvailable bool     `json:"minQuantityAvailable"`
	AvailableQuantity    int32    `json:"availableQuantity"`
	Solutions            []string `json:"solutions"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *LocationAvailabilityDetail) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
