package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// GetPrivateNetworkCommand represents command that retrieves details about speific private network on the account
type GetPrivateNetworkCommand struct {
	receiver  receiver.BMCSDK
	networkID string
}

// Execute runs GetPrivateNetworkCommand
func (command *GetPrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksNetworkIdGet(context.Background(), command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("GetPrivateNetworkCommand %s", errResolver.Error)
}

//NewGetPrivateNetworkCommand constructs new commmand of this type
func NewGetPrivateNetworkCommand(receiver receiver.BMCSDK, networkID string) *GetPrivateNetworkCommand {

	return &GetPrivateNetworkCommand{receiver, networkID}
}
