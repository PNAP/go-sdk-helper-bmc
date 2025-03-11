package publicnetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
)

// GetPublicNetworksCommand represents command that lists all public networks on the account
type GetPublicNetworksCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetPublicNetworksCommand
func (command *GetPublicNetworksCommand) Execute() ([]networkapiclient.PublicNetwork, error) {

	networks, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksAPI.PublicNetworksGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return networks, nil
	}
	return nil, fmt.Errorf("GetPublicNetworksCommand %s", errResolver.Error)
}

//NewGetPublicNetworksCommand constructs new commmand of this type
func NewGetPublicNetworksCommand(receiver receiver.BMCSDK) *GetPublicNetworksCommand {

	return &GetPublicNetworksCommand{receiver}
}
