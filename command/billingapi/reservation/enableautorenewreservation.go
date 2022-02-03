package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// EnableAutoRenewReservationCommand represents command that enables auto-renewal for specific reservation.
// Use NewEnableAutoRenewReservationCommand to initialize command properly.
type EnableAutoRenewReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
}

// Execute enables auto-renewal for specific reservation.
func (command *EnableAutoRenewReservationCommand) Execute() (*dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"

	httpResponse, err := req.PNAPClient.Post(apiPrefix+"reservations/"+command.reservationID+"/actions/auto-renew/enable", nil)

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = dto.Reservation{}
		reservationResponse.FromBytes(httpResponse)
		return &reservationResponse, nil
	}
	return nil, fmt.Errorf("EnableAutoRenewReservationCommand %s", errResolver.Error)

}

//NewEnableAutoRenewReservationCommand constructs new commmand of this type
func NewEnableAutoRenewReservationCommand(requester receiver.BMCSDK, reservationID string) *EnableAutoRenewReservationCommand {

	return &EnableAutoRenewReservationCommand{requester, reservationID}
}
