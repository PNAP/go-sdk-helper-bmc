package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
	"time"
)

//Reservation
type Reservation struct {
	ID                  string    `json:"id"`
	ProductCode         string    `json:"productCode"`
	ProductCategory     string    `json:"productCategory"`
	Location            string    `json:"location"`
	ReservationModel    string    `json:"reservationModel"`
	InitialInvoiceModel string    `json:"initialInvoiceModel,omitempty"`
	StartDateTime       time.Time `json:"startDateTime"`
	EndDateTime         time.Time `json:"endDateTime,omitempty"`
	LastRenewalDateTime time.Time `json:"lastRenewalDateTime,omitempty"`
	NextRenewalDateTime time.Time `json:"nextRenewalDateTime,omitempty"`
	AutoRenew           bool      `json:"autoRenew"`
	SKU                 string    `json:"sku"`
	Price               float32   `json:"price"`
	PriceUnit           string    `json:"priceUnit"`
	AssignedResourceID  string    `json:"assignedResourceId,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *Reservation) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
