package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//Reservation
type Reservation struct {
	SKU string `json:"sku,omitempty"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto Reservation) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("Reservation dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}

//FromBytes performs conversion of http response to the representing struct
func (dto *Reservation) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
