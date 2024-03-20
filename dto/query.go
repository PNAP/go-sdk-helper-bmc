package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"

	locationapiclient "github.com/phoenixnap/go-sdk-bmc/locationapi/v3"
)

//Query represents query parameters used in various API calls
type Query struct {
	Force           bool                                  `json:"force"`
	Location        locationapiclient.LocationEnum        `json:"location"`
	ProductCategory locationapiclient.ProductCategoryEnum `json:"productCategory"`
	Number          string                                `json:"number"`
	Status          string                                `json:"status"`
	SentOnFrom      time.Time                             `json:"sentOnFrom"`
	SentOnTo        time.Time                             `json:"sentOnTo"`
	Limit           int32                                 `json:"limit,omitempty"`
	Offset          int32                                 `json:"offset"`
	SortField       string                                `json:"sortField"`
	SortDirection   string                                `json:"sortDirection"`
	From            time.Time                             `json:"from"`
	To              time.Time                             `json:"to"`
	Order           string                                `json:"order"`
	Username        string                                `json:"username"`
	Verb            string                                `json:"verb"`
	Uri             string                                `json:"uri"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto Query) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("Query dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
