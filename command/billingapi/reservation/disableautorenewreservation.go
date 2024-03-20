package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// DisableAutoRenewReservationCommand represents command that disables auto-renewal for specific reservation.
// Use NewDisableAutoRenewReservationCommand to initialize command properly.
type DisableAutoRenewReservationCommand struct {
	receiver      receiver.BMCSDK
	reservationID string
	reason        billingapiclient.ReservationAutoRenewDisableRequest
}

// Execute disables auto-renewal for specific reservation.
func (command *DisableAutoRenewReservationCommand) Execute() (*billingapiclient.Reservation, error) {

	reservation, httpResponse, err := command.receiver.BillingAPIClient.ReservationsAPI.ReservationsReservationIdActionsAutoRenewDisablePost(context.Background(), command.reservationID).
		ReservationAutoRenewDisableRequest(command.reason).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservation, nil
	}
	return nil, fmt.Errorf("DisableAutoRenewReservationCommand %s", errResolver.Error)
}

//NewDisableAutoRenewReservationCommand constructs new commmand of this type
func NewDisableAutoRenewReservationCommand(receiver receiver.BMCSDK, reservationID string, reason billingapiclient.ReservationAutoRenewDisableRequest) *DisableAutoRenewReservationCommand {

	return &DisableAutoRenewReservationCommand{receiver, reservationID, reason}
}
