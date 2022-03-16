package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

//ProductQuery represents query parameters for products retrieval
type ProductQuery struct {
	ProductCode     string `json:"productCode,omitempty"`
	ProductCategory string `json:"productCategory,omitempty"`
	SKUCode         string `json:"skucode,omitempty"`
	Location        string `json:"location,omitempty"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto ProductQuery) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("ProductQuery dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
