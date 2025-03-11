package publicnetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
)

// UpdatePublicNetworkCommand represents command that updates public network on the server
type UpdatePublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	networkID           string
	publicNetworkModify networkapiclient.PublicNetworkModify
}

// Execute runs UpdatePublicNetworkCommand
func (command *UpdatePublicNetworkCommand) Execute() (*networkapiclient.PublicNetwork, error) {

	network, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksAPI.PublicNetworksNetworkIdPatch(context.Background(), command.networkID).PublicNetworkModify(command.publicNetworkModify).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return network, nil
	}
	return nil, fmt.Errorf("UpdatePublicNetworkCommand %s", errResolver.Error)
}

//NewUpdatePublicNetworkCommand constructs new commmand of this type
func NewUpdatePublicNetworkCommand(receiver receiver.BMCSDK, networkID string, publicNetworkModify networkapiclient.PublicNetworkModify) *UpdatePublicNetworkCommand {

	return &UpdatePublicNetworkCommand{receiver, networkID, publicNetworkModify}
}
