package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
)

// GetPrivateNetworksCommand represents command that lists all private networks on the account
type GetPrivateNetworksCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetPrivateNetworksCommand
func (command *GetPrivateNetworksCommand) Execute() ([]networkapiclient.PrivateNetwork, error) {

	servers, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return servers, nil
	}
	return nil, fmt.Errorf("GetPrivateNetworksCommand %s", errResolver.Error)
}

//NewGetPrivateNetworksCommand constructs new commmand of this type
func NewGetPrivateNetworksCommand(receiver receiver.BMCSDK) *GetPrivateNetworksCommand {

	return &GetPrivateNetworksCommand{receiver}
}
