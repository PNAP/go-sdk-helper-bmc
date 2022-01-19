package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

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
