package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// EnableAutoRenewReservationCommand represents command that enables auto-renewal for specific reservation.
// Use NewEnableAutoRenewReservationCommand to initialize command properly.
type EnableAutoRenewReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
}

// Execute enables auto-renewal for specific reservation.
func (command *EnableAutoRenewReservationCommand) Execute() (*billingapiclient.Reservation, error) {

	reservation, httpResponse, err := command.receiver.BillingAPIClient.ReservationsAPI.ReservationsReservationIdActionsAutoRenewEnablePost(context.Background(), command.reservationID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservation, nil
	}
	return nil, fmt.Errorf("EnableAutoRenewReservationCommand %s", errResolver.Error)
}

//NewEnableAutoRenewReservationCommand constructs new commmand of this type
func NewEnableAutoRenewReservationCommand(receiver receiver.BMCSDK, reservationID string) *EnableAutoRenewReservationCommand {

	return &EnableAutoRenewReservationCommand{receiver, reservationID}
}
