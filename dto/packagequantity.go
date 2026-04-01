package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

// PackageQuantity represents the package size per month, for a product's pricing plan,
// including the min quantity and max quantity information.
type PackageQuantity struct {
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}

// FromBytes performs conversion of http response to the representing struct
func (dto *PackageQuantity) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
