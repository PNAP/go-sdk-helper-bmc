package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// GetReservationCommand represents command that retrieves specified reservation for the account.
// Use NewGetReservationCommand to initialize command properly.
type GetReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
}

// Execute retrieves specified reservation for the account.
func (command *GetReservationCommand) Execute() (*billingapiclient.Reservation, error) {

	reservation, httpResponse, err := command.receiver.BillingAPIClient.ReservationsAPI.ReservationsReservationIdGet(context.Background(), command.reservationID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservation, nil
	}
	return nil, fmt.Errorf("GetReservationCommand %s", errResolver.Error)
}

//NewGetReservationCommand constructs new commmand of this type
func NewGetReservationCommand(receiver receiver.BMCSDK, reservationID string) *GetReservationCommand {

	return &GetReservationCommand{receiver, reservationID}
}
