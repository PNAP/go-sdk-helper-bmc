package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//Reservations represents list of reservations
type Reservations []Reservation

//FromBytes performs conversion of http response to the representing struct
func (dto *Reservations) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
