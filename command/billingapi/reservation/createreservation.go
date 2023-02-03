package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi"
)

// CreateReservationCommand represents command that creates new package reservation for the account.
// Use NewCreateReservationCommand to initialize command properly.
type CreateReservationCommand struct {
	receiver           receiver.BMCSDK
	reservationRequest billingapiclient.ReservationRequest
}

// Execute creates new package reservation for the account.
func (command *CreateReservationCommand) Execute() (*billingapiclient.Reservation, error) {

	reservation, httpResponse, err := command.receiver.BillingAPIClient.ReservationsApi.ReservationsPost(context.Background()).ReservationRequest(command.reservationRequest).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservation, nil
	}
	return nil, fmt.Errorf("CreateReservationCommand %s", errResolver.Error)
}

//NewCreateReservationCommand constructs new commmand of this type
func NewCreateReservationCommand(receiver receiver.BMCSDK, reservationRequest billingapiclient.ReservationRequest) *CreateReservationCommand {

	return &CreateReservationCommand{receiver, reservationRequest}
}
