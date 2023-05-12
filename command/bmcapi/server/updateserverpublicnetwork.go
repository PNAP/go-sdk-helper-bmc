package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// UpdateServerPublicNetworkCommand represents command that updates public network ip adresses on specific server
type UpdateServerPublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	publicNetworkID     string
	serverNetworkUpdate bmcapiclient.ServerNetworkUpdate
	query               dto.Query
}

// Execute runs UpdateServerPublicNetworkCommand
func (command *UpdateServerPublicNetworkCommand) Execute() (*bmcapiclient.ServerPublicNetwork, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPublicNetworksPatch(context.Background(), command.serverID, command.publicNetworkID).Force(force).ServerNetworkUpdate(command.serverNetworkUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("UpdateServerPublicNetworkCommand %s", errResolver.Error)
}

//NewUpdateServerPublicNetworkCommand constructs new commmand of this type
func NewUpdateServerPublicNetworkCommand(receiver receiver.BMCSDK, serverID string, publicNetworkID string, serverNetworkUpdate bmcapiclient.ServerNetworkUpdate, query dto.Query) *UpdateServerPublicNetworkCommand {

	return &UpdateServerPublicNetworkCommand{receiver, serverID, publicNetworkID, serverNetworkUpdate, query}
}
