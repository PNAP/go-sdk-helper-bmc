package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

// PackageDetails represents details object which contains package quantity and its unit.
type PackageDetails struct {
	PackageQuantity *PackageQuantity `json:"packageQuantity,omitempty"`
	PackageUnit     *string          `json:"packageUnit,omitempty"`
}

// FromBytes performs conversion of http response to the representing struct
func (dto *PackageDetails) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
