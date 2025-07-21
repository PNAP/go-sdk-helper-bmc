package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// TransferServerReservationCommand represents command that transfers reservation from a specific server to another server
type TransferServerReservationCommand struct {
	receiver                   receiver.BMCSDK
	serverID                   string
	reservationTransferDetails bmcapiclient.ReservationTransferDetails
}

// Execute transfers reservation from a specific server to another server
func (command *TransferServerReservationCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdActionsTransferReservation(context.Background(), command.serverID).
		ReservationTransferDetails(command.reservationTransferDetails).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("TransferServerReservationCommand %s", errResolver.Error)
}

// NewTransferServerReservationCommand constructs new commmand of this type
func NewTransferServerReservationCommand(receiver receiver.BMCSDK, serverID string, reservationTransferDetails bmcapiclient.ReservationTransferDetails) *TransferServerReservationCommand {

	return &TransferServerReservationCommand{receiver, serverID, reservationTransferDetails}
}
