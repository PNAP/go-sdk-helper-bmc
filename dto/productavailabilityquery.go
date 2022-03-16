package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

//ProductAvailabilityQuery represents query parameters for product availabilities retrieval
type ProductAvailabilityQuery struct {
	ProductCategory              *[]string `json:"productCategory,omitempty"`
	ProductCode                  *[]string `json:"productCode,omitempty"`
	ShowOnlyMinQuantityAvailable *bool     `json:"showOnlyMinQuantityAvailable,omitempty"`
	Location                     *[]string `json:"location,omitempty"`
	Solution                     *[]string `json:"solution,omitempty"`
	MinQuantity                  *int32    `json:"minQuantity,omitempty"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto ProductAvailabilityQuery) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("ProductAvailabilityQuery dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
