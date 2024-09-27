package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//GPUConfigurationMetadata represents details of the GPU configuration
type GPUConfigurationMetadata struct {
	Count float32 `json:"count,omitempty"`
	Name  string  `json:"name,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *GPUConfigurationMetadata) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
