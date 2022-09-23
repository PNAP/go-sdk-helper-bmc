package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
)

// UpdatePrivateNetworkCommand represents command that updates private network on the server
type UpdatePrivateNetworkCommand struct {
	receiver             receiver.BMCSDK
	networkID            string
	privateNetworkModify networkapiclient.PrivateNetworkModify
}

// Execute runs UpdatePrivateNetworkCommand
func (command *UpdatePrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksNetworkIdPut(context.Background(), command.networkID).PrivateNetworkModify(command.privateNetworkModify).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("UpdatePrivateNetworkCommand %s", errResolver.Error)
}

//NewUpdatePrivateNetworkCommand constructs new commmand of this type
func NewUpdatePrivateNetworkCommand(receiver receiver.BMCSDK, networkID string, privateNetworkModify networkapiclient.PrivateNetworkModify) *UpdatePrivateNetworkCommand {

	return &UpdatePrivateNetworkCommand{receiver, networkID, privateNetworkModify}
}
