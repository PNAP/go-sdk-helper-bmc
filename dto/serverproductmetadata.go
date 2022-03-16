package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//ServerProductMetadata represents details of the server product
type ServerProductMetadata struct {
	RamInGb      int32   `json:"ramInGb"`
	CPU          string  `json:"cpu"`
	CPUCount     int32   `json:"cpuCount"`
	CoresPerCPU  int32   `json:"coresPerCpu"`
	CPUFrequency float32 `json:"cpuFrequency"`
	Network      string  `json:"network"`
	Storage      string  `json:"storage"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ServerProductMetadata) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
