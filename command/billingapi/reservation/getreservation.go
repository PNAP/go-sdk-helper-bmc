package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// GetReservationCommand represents command that retrieves specified reservation for the account.
// Use NewGetReservationCommand to initialize command properly.
type GetReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
}

// Execute retrieves specified reservation for the account.
func (command *GetReservationCommand) Execute() (*dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"

	httpResponse, err := req.PNAPClient.Get(apiPrefix + "reservations/" + command.reservationID)

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = dto.Reservation{}
		reservationResponse.FromBytes(httpResponse)
		return &reservationResponse, nil
	}
	return nil, fmt.Errorf("GetReservationCommand %s", errResolver.Error)

}

//NewGetReservationCommand constructs new commmand of this type
func NewGetReservationCommand(requester receiver.BMCSDK, reservationID string) *GetReservationCommand {

	return &GetReservationCommand{requester, reservationID}
}
