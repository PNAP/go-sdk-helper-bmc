package server

import (
	"fmt"
	"net/http"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// AddServer2PrivateNetworkCommand represents command that configures private network on specific server
type AddServer2PrivateNetworkCommand struct {
	receiver             receiver.BMCSDK
	serverID             string
	serverPrivateNetwork bmcapiclient.ServerPrivateNetwork
	query                *dto.Query
}

// Execute runs AddServer2PrivateNetworkCommand
func (command *AddServer2PrivateNetworkCommand) Execute() (*bmcapiclient.ServerPrivateNetwork, error) {

	var server *bmcapiclient.ServerPrivateNetwork
	var httpResponse *http.Response
	var err error

	if command.query != nil {

		force := command.query.Force

		server, httpResponse, err = command.receiver.APIClient.ServersAPI.ServersServerIdPrivateNetworksPost(context.Background(), command.serverID).Force(force).
			ServerPrivateNetwork(command.serverPrivateNetwork).Execute()
	} else {

		server, httpResponse, err = command.receiver.APIClient.ServersAPI.ServersServerIdPrivateNetworksPost(context.Background(), command.serverID).
			ServerPrivateNetwork(command.serverPrivateNetwork).Execute()
	}

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("AddServer2PrivateNetworkCommand %s", errResolver.Error)
}

//NewAddServer2PrivateNetworkCommand constructs new commmand of this type
func NewAddServer2PrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPrivateNetwork bmcapiclient.ServerPrivateNetwork) *AddServer2PrivateNetworkCommand {

	return &AddServer2PrivateNetworkCommand{receiver, serverID, serverPrivateNetwork, nil}
}

//NewAddServer2PrivateNetworkCommandWithQuery constructs new commmand of this type
func NewAddServer2PrivateNetworkCommandWithQuery(receiver receiver.BMCSDK, serverID string, serverPrivateNetwork bmcapiclient.ServerPrivateNetwork, query *dto.Query) *AddServer2PrivateNetworkCommand {

	return &AddServer2PrivateNetworkCommand{receiver, serverID, serverPrivateNetwork, query}
}
