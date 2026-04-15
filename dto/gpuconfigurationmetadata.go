package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

// GpuConfigurationMetadata represents details of the GPU configuration
type GpuConfigurationMetadata struct {
	Count *float32 `json:"count,omitempty"`
	Name  *string  `json:"name,omitempty"`
}

// FromBytes performs conversion of http response to the representing struct
func (dto *GpuConfigurationMetadata) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
