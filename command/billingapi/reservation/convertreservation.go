package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// ConvertReservationCommand represents command that converts specified reservation for the account.
// Use NewConvertReservationCommand to initialize command properly.
type ConvertReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
	reservation   dto.ReservationRequest
}

// Execute converts specified reservation for the account.
func (command *ConvertReservationCommand) Execute() (*dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"
	val, err := command.reservation.ToBytes()
	if err != nil {
		return nil, err
	}

	httpResponse, err := req.PNAPClient.Post(apiPrefix+"reservations/"+command.reservationID+"/actions/convert", val)

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = dto.Reservation{}
		reservationResponse.FromBytes(httpResponse)
		return &reservationResponse, nil
	}
	return nil, fmt.Errorf("ConvertReservationCommand %s", errResolver.Error)

}

//NewConvertReservationCommand constructs new commmand of this type
func NewConvertReservationCommand(requester receiver.BMCSDK, reservationID string, reservation dto.ReservationRequest) *ConvertReservationCommand {

	return &ConvertReservationCommand{requester, reservationID, reservation}
}
