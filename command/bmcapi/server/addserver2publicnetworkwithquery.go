package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// AddServer2PublicNetworkWithQueryCommand represents command that configures public network on specific server with optional use of query parameter
type AddServer2PublicNetworkWithQueryCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	serverPublicNetwork bmcapiclient.ServerPublicNetwork
	query               dto.Query
}

// Execute runs AddServer2PublicNetworkWithQueryCommand
func (command *AddServer2PublicNetworkWithQueryCommand) Execute() (*bmcapiclient.ServerPublicNetwork, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdPublicNetworksPost(context.Background(), command.serverID).Force(force).
		ServerPublicNetwork(command.serverPublicNetwork).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("AddServer2PublicNetworkWithQueryCommand %s", errResolver.Error)
}

//NewAddServer2PublicNetworkWithQueryCommand constructs new commmand of this type
func NewAddServer2PublicNetworkWithQueryCommand(receiver receiver.BMCSDK, serverID string, serverPublicNetwork bmcapiclient.ServerPublicNetwork, query dto.Query) *AddServer2PublicNetworkWithQueryCommand {

	return &AddServer2PublicNetworkWithQueryCommand{receiver, serverID, serverPublicNetwork, query}
}
