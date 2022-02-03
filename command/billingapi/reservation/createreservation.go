package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// CreateReservationCommand represents command that creates new package reservation for the account.
// Use NewCreateReservationCommand to initialize command properly.
type CreateReservationCommand struct {
	receiver    receiver.BMCSDK
	reservation dto.ReservationRequest
}

// Execute creates new package reservation for the account.
func (command *CreateReservationCommand) Execute() (*dto.Reservation, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/"
	val, err := command.reservation.ToBytes()
	if err != nil {
		return nil, err
	}

	httpResponse, err := req.PNAPClient.Post(apiPrefix+"reservations", val)

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var reservationResponse = dto.Reservation{}
		reservationResponse.FromBytes(httpResponse)
		return &reservationResponse, nil
	}
	return nil, fmt.Errorf("CreateReservationCommand %s", errResolver.Error)

}

//NewCreateReservationCommand constructs new commmand of this type
func NewCreateReservationCommand(requester receiver.BMCSDK, reservation dto.ReservationRequest) *CreateReservationCommand {

	return &CreateReservationCommand{requester, reservation}
}
