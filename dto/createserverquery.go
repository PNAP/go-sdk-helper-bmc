package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

//CreateServerQuery represents query parameters for server provisioning
type CreateServerQuery struct {
	Force bool `json:"force"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto CreateServerQuery) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("CreateServerQuery dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
