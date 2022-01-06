package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// CreateServerPrivateNetworkCommand represents command that configures private network on specific server
type CreateServerPrivateNetworkCommand struct {
	receiver             receiver.BMCSDK
	serverID             string
	serverPrivateNetwork bmcapiclient.ServerPrivateNetwork
}

// Execute con command on specific server
func (command *CreateServerPrivateNetworkCommand) Execute() (*bmcapiclient.ServerPrivateNetwork, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPrivateNetworksPost(context.Background(), command.serverID).ServerPrivateNetwork(command.serverPrivateNetwork).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &server, nil
	}
	return nil, fmt.Errorf("CreateServerPrivateNetworkCommand %s", errResolver.Error)
}

//NewCreateServerPrivateNetworkCommand constructs new commmand of this type
func NewCreateServerPrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPrivateNetwork bmcapiclient.ServerPrivateNetwork) *CreateServerPrivateNetworkCommand {

	return &CreateServerPrivateNetworkCommand{receiver, serverID, serverPrivateNetwork}
}
