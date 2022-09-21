package publicnetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// GetPublicNetworkCommand represents command that retrieves details about speific public network on the account
type GetPublicNetworkCommand struct {
	receiver  receiver.BMCSDK
	networkID string
}

// Execute runs GetPublicNetworkCommand
func (command *GetPublicNetworkCommand) Execute() (*networkapiclient.PublicNetwork, error) {

	network, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksApi.PublicNetworksNetworkIdGet(context.Background(), command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return network, nil
	}
	return nil, fmt.Errorf("GetPublicNetworkCommand %s", errResolver.Error)
}

//NewGetPublicNetworkCommand constructs new commmand of this type
func NewGetPublicNetworkCommand(receiver receiver.BMCSDK, networkID string) *GetPublicNetworkCommand {

	return &GetPublicNetworkCommand{receiver, networkID}
}
