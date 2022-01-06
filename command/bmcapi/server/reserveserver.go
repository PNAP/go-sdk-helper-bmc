package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// ReserveServerCommand represents command that reserves specific server
type ReserveServerCommand struct {
	receiver      receiver.BMCSDK
	serverID      string
	serverReserve bmcapiclient.ServerReserve
}

// Execute reserves specific server
func (command *ReserveServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsReservePost(context.Background(), command.serverID).ServerReserve(command.serverReserve).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &server, nil
	}
	return nil, fmt.Errorf("ReserveServerCommand %s", errResolver.Error)
}

//NewReserveServerCommand constructs new commmand of this type
func NewReserveServerCommand(receiver receiver.BMCSDK, serverID string, serverReserve bmcapiclient.ServerReserve) *ReserveServerCommand {

	return &ReserveServerCommand{receiver, serverID, serverReserve}
}
