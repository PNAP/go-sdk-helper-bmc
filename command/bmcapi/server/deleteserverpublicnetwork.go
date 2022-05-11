package server

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeleteServerPublicNetworkCommand represents command that removes public network from networking devices for specific server
type DeleteServerPublicNetworkCommand struct {
	receiver  receiver.BMCSDK
	serverID  string
	networkID string
}

// Execute runs DeleteServerPublicNetworkCommand for specific server
func (command *DeleteServerPublicNetworkCommand) Execute() (*string, error) {

	response, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPublicNetworksDelete(context.Background(), command.serverID, command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &response, nil
	}
	return nil, fmt.Errorf("DeleteServerPublicNetworkCommand %s", errResolver.Error)
}

//NewDeleteServerPublicNetworkCommand constructs new commmand of this type
func NewDeleteServerPublicNetworkCommand(receiver receiver.BMCSDK, serverID string, networkID string) *DeleteServerPublicNetworkCommand {

	return &DeleteServerPublicNetworkCommand{receiver, serverID, networkID}
}
