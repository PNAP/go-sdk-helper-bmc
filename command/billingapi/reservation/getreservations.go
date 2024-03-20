package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// GetReservationsCommand represents command that retrieves all reservations for the account.
// Use NewGetReservationsCommand to initialize command properly.
type GetReservationsCommand struct {
	receiver receiver.BMCSDK
}

// Execute retrieves all reservations for the account.
func (command *GetReservationsCommand) Execute() ([]billingapiclient.Reservation, error) {

	reservations, httpResponse, err := command.receiver.BillingAPIClient.ReservationsAPI.ReservationsGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservations, nil
	}
	return nil, fmt.Errorf("GetReservationsCommand %s", errResolver.Error)
}

//NewGetReservationsCommand constructs new commmand of this type
func NewGetReservationsCommand(receiver receiver.BMCSDK) *GetReservationsCommand {

	return &GetReservationsCommand{receiver}
}
