package reservation

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// ConvertReservationCommand represents command that converts specified reservation for the account.
// Use NewConvertReservationCommand to initialize command properly.
type ConvertReservationCommand struct {
	receiver           receiver.BMCSDK
	reservationID      string
	reservationRequest billingapiclient.ReservationRequest
}

// Execute converts specified reservation for the account.
func (command *ConvertReservationCommand) Execute() (*billingapiclient.Reservation, error) {

	reservation, httpResponse, err := command.receiver.BillingAPIClient.ReservationsAPI.ReservationsReservationIdActionsConvertPost(context.Background(), command.reservationID).
		ReservationRequest(command.reservationRequest).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return reservation, nil
	}
	return nil, fmt.Errorf("ConvertReservationCommand %s", errResolver.Error)
}

//NewConvertReservationCommand constructs new commmand of this type
func NewConvertReservationCommand(receiver receiver.BMCSDK, reservationID string, reservationRequest billingapiclient.ReservationRequest) *ConvertReservationCommand {

	return &ConvertReservationCommand{receiver, reservationID, reservationRequest}
}
