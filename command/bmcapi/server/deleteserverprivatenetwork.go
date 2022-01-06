package server

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeleteServerPrivateNetworkCommand represents command that removes private network from networking devices for specific server
type DeleteServerPrivateNetworkCommand struct {
	receiver  receiver.BMCSDK
	serverID  string
	networkID string
}

// Execute runs DeleteServerPrivateNetwork command for specific server
func (command *DeleteServerPrivateNetworkCommand) Execute() (*string, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.DeletePrivateNetwork(context.Background(), command.serverID, command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &server, nil
	}
	return nil, fmt.Errorf("DeleteServerPrivateNetworkCommand %s", errResolver.Error)
}

//NewDeleteServerPrivateNetworkCommand constructs new commmand of this type
func NewDeleteServerPrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, networkID string) *DeleteServerPrivateNetworkCommand {

	return &DeleteServerPrivateNetworkCommand{receiver, serverID, networkID}
}
