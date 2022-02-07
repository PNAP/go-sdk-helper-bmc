package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// GetReservationsCommand represents command that retrieves all reservations for the account.
// Use NewGetReservationsCommand to initialize command properly.
type GetReservationsCommand struct {
	receiver receiver.BMCSDK
}

// Execute retrieves all reservations for the account.
func (command *GetReservationsCommand) Execute() ([]dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"

	httpResponse, err := req.PNAPClient.Get(apiPrefix + "reservations")

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = &dto.Reservations{}
		reservationResponse.FromBytes(httpResponse)
		respList := *reservationResponse
		return respList, nil
	}
	return nil, fmt.Errorf("GetReservationsCommand %s", errResolver.Error)

}

//NewGetReservationsCommand constructs new commmand of this type
func NewGetReservationsCommand(requester receiver.BMCSDK) *GetReservationsCommand {

	return &GetReservationsCommand{requester}
}
