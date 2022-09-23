package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// AddServer2PublicNetworkCommand represents command that configures public network on specific server
type AddServer2PublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	serverPublicNetwork bmcapiclient.ServerPublicNetwork
}

// Execute runs AddServer2PublicNetworkCommand
func (command *AddServer2PublicNetworkCommand) Execute() (*bmcapiclient.ServerPublicNetwork, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPublicNetworksPost(context.Background(), command.serverID).ServerPublicNetwork(command.serverPublicNetwork).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("AddServer2PublicNetworkCommand %s", errResolver.Error)
}

//NewAddServer2PublicNetworkCommand constructs new commmand of this type
func NewAddServer2PublicNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPublicNetwork bmcapiclient.ServerPublicNetwork) *AddServer2PublicNetworkCommand {

	return &AddServer2PublicNetworkCommand{receiver, serverID, serverPublicNetwork}
}
