package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

//Query represents query parameters used in various API calls
type Query struct {
	Force bool `json:"force"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto Query) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("CreateServerQuery dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
