package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// AddServer2PrivateNetworkWithQueryCommand represents command that configures private network on specific server with optional use of query parameter
type AddServer2PrivateNetworkWithQueryCommand struct {
	receiver             receiver.BMCSDK
	serverID             string
	serverPrivateNetwork bmcapiclient.ServerPrivateNetwork
	query                dto.Query
}

// Execute runs AddServer2PrivateNetworkWithQueryCommand
func (command *AddServer2PrivateNetworkWithQueryCommand) Execute() (*bmcapiclient.ServerPrivateNetwork, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdPrivateNetworksPost(context.Background(), command.serverID).Force(force).
		ServerPrivateNetwork(command.serverPrivateNetwork).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("AddServer2PrivateNetworkWithQueryCommand %s", errResolver.Error)
}

//NewAddServer2PrivateNetworkWithQueryCommand constructs new commmand of this type
func NewAddServer2PrivateNetworkWithQueryCommand(receiver receiver.BMCSDK, serverID string, serverPrivateNetwork bmcapiclient.ServerPrivateNetwork, query dto.Query) *AddServer2PrivateNetworkWithQueryCommand {

	return &AddServer2PrivateNetworkWithQueryCommand{receiver, serverID, serverPrivateNetwork, query}
}
