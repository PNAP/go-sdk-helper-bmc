package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

// ServerProductMetadata represents details of the server product
type ServerProductMetadata struct {
	RamInGb           float32                    `json:"ramInGb"`
	Cpu               string                     `json:"cpu"`
	CpuCount          float32                    `json:"cpuCount"`
	CoresPerCpu       float32                    `json:"coresPerCpu"`
	CpuFrequency      float32                    `json:"cpuFrequency"`
	Network           string                     `json:"network"`
	Storage           string                     `json:"storage"`
	GpuConfigurations []GpuConfigurationMetadata `json:"gpuConfigurations,omitempty"`
}

// FromBytes performs conversion of http response to the representing struct
func (dto *ServerProductMetadata) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
