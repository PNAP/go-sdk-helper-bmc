package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//Reservation request
type ReservationRequest struct {
	SKU string `json:"sku"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto ReservationRequest) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("ReservationRequest dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}

//FromBytes performs conversion of http response to the representing struct
func (dto *ReservationRequest) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}

//Reservation auto-renewal disabling request reason
type ReservationAutoRenewDisableRequest struct {
	AutoRenewDisableReason *string `json:"autoRenewDisableReason,omitempty"`
}

//ToBytes performs conversion of struct to the io.Reader
func (dto ReservationAutoRenewDisableRequest) ToBytes() (io.Reader, error) {
	requestByte, err := json.Marshal(dto)

	if err != nil {
		return nil, fmt.Errorf("ReservationAutoRenewDisableRequest dto can not be converted to io.Reader: %s", err)
	}
	return bytes.NewBuffer(requestByte), nil
}
