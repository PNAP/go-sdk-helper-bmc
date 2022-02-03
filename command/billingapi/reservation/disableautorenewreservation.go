package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DisableAutoRenewReservationCommand represents command that disables auto-renewal for specific reservation.
// Use NewDisableAutoRenewReservationCommand to initialize command properly.
type DisableAutoRenewReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
	reason        dto.ReservationAutoRenewDisableRequest
}

// Execute disables auto-renewal for specific reservation.
func (command *DisableAutoRenewReservationCommand) Execute() (*dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"
	val, err := command.reason.ToBytes()
	if err != nil {
		return nil, err
	}

	httpResponse, err := req.PNAPClient.Post(apiPrefix+"reservations/"+command.reservationID+"/actions/auto-renew/disable", val)

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = dto.Reservation{}
		reservationResponse.FromBytes(httpResponse)
		return &reservationResponse, nil
	}
	return nil, fmt.Errorf("DisableAutoRenewReservationCommand %s", errResolver.Error)

}

//NewDisableAutoRenewReservationCommand constructs new commmand of this type
func NewDisableAutoRenewReservationCommand(requester receiver.BMCSDK, reservationID string, reason dto.ReservationAutoRenewDisableRequest) *DisableAutoRenewReservationCommand {

	return &DisableAutoRenewReservationCommand{requester, reservationID, reason}
}
