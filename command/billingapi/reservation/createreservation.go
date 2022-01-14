package reservation

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// CreateReservationCommand represents command for server provisioning. Use NewCreateServerCommand to initilize command properly.
type CreateReservationCommand struct {
	receiver    receiver.BMCSDK
	reservation dto.Reservation
}

// Execute provisions new server
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

//NewCreateServerCommand constructs new commmand of this type
func NewCreateReservationCommand(requester receiver.BMCSDK, server dto.Reservation) *CreateReservationCommand {

	return &CreateReservationCommand{requester, server}
}
