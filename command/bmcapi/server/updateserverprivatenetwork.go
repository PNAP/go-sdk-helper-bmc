package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// UpdateServerPrivateNetworkCommand represents command that updates private network ip adresses on specific server
type UpdateServerPrivateNetworkCommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	privateNetworkID    string
	serverNetworkUpdate bmcapiclient.ServerNetworkUpdate
	query               dto.Query
}

// Execute runs UpdateServerPrivateNetworkCommand
func (command *UpdateServerPrivateNetworkCommand) Execute() (*bmcapiclient.ServerPrivateNetwork, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPrivateNetworksPatch(context.Background(), command.serverID, command.privateNetworkID).Force(force).ServerNetworkUpdate(command.serverNetworkUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("UpdateServerPrivateNetworkCommand %s", errResolver.Error)
}

//NewUpdateServerPrivateNetworkCommand constructs new commmand of this type
func NewUpdateServerPrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, privateNetworkID string, serverNetworkUpdate bmcapiclient.ServerNetworkUpdate, query dto.Query) *UpdateServerPrivateNetworkCommand {

	return &UpdateServerPrivateNetworkCommand{receiver, serverID, privateNetworkID, serverNetworkUpdate, query}
}
