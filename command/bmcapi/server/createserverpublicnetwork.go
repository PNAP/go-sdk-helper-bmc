package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// CreateServerPublicNetworkCommand represents command that configures public network on specific server
type CreateServerPublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	serverPublicNetwork bmcapiclient.ServerPublicNetwork
}

// Execute runs CreateServerPublicNetworkCommand on specific server
func (command *CreateServerPublicNetworkCommand) Execute() (*bmcapiclient.ServerPublicNetwork, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPublicNetworksPost(context.Background(), command.serverID).ServerPublicNetwork(command.serverPublicNetwork).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &server, nil
	}
	return nil, fmt.Errorf("CreateServerPublicNetworkCommand %s", errResolver.Error)
}

//NewCreateServerPublicNetworkCommand constructs new commmand of this type
func NewCreateServerPublicNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPublicNetwork bmcapiclient.ServerPublicNetwork) *CreateServerPublicNetworkCommand {

	return &CreateServerPublicNetworkCommand{receiver, serverID, serverPublicNetwork}
}
