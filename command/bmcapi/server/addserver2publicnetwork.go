package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// AddServer2PublicNetworkCommand represents command that configures public network on specific server
type AddServer2PublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	serverPublicNetwork bmcapiclient.ServerPublicNetwork
	query               *dto.Query
}

// Execute runs AddServer2PublicNetworkCommand
func (command *AddServer2PublicNetworkCommand) Execute() (*bmcapiclient.ServerPublicNetwork, error) {

	if command.query != nil {

		force := command.query.Force

		server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdPublicNetworksPost(context.Background(), command.serverID).Force(force).
			ServerPublicNetwork(command.serverPublicNetwork).Execute()

		errResolver := dto.NewErrorResolver(httpResponse, err)

		if errResolver.Error == nil {
			return server, nil
		}
		return nil, fmt.Errorf("AddServer2PublicNetworkWithQueryCommand %s", errResolver.Error)

	} else {

		server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdPublicNetworksPost(context.Background(), command.serverID).
			ServerPublicNetwork(command.serverPublicNetwork).Execute()

		errResolver := dto.NewErrorResolver(httpResponse, err)

		if errResolver.Error == nil {
			return server, nil
		}
		return nil, fmt.Errorf("AddServer2PublicNetworkCommand %s", errResolver.Error)

	}
}

//NewAddServer2PublicNetworkCommand constructs new commmand of this type
func NewAddServer2PublicNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPublicNetwork bmcapiclient.ServerPublicNetwork) *AddServer2PublicNetworkCommand {

	return &AddServer2PublicNetworkCommand{receiver, serverID, serverPublicNetwork, nil}
}

//NewAddServer2PublicNetworkCommandWithQuery constructs new commmand of this type
func NewAddServer2PublicNetworkCommandWithQuery(receiver receiver.BMCSDK, serverID string, serverPublicNetwork bmcapiclient.ServerPublicNetwork, query *dto.Query) *AddServer2PublicNetworkCommand {

	return &AddServer2PublicNetworkCommand{receiver, serverID, serverPublicNetwork, query}
}
