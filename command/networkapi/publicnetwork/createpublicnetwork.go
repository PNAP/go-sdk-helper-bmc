package publicnetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// CreatePublicNetworkCommand represents command that creates new public network
type CreatePublicNetworkCommand struct {
	receiver            receiver.BMCSDK
	publicNetworkCreate networkapiclient.PublicNetworkCreate
}

// Execute runs CreatePublicNetworkCommand
func (command *CreatePublicNetworkCommand) Execute() (*networkapiclient.PublicNetwork, error) {

	network, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksApi.PublicNetworksPost(context.Background()).PublicNetworkCreate(command.publicNetworkCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return network, nil
	}
	return nil, fmt.Errorf("CreatePublicNetworkCommand %s", errResolver.Error)
}

//NewCreatePublicNetworkCommand constructs new commmand of this type
func NewCreatePublicNetworkCommand(receiver receiver.BMCSDK, publicNetworkCreate networkapiclient.PublicNetworkCreate) *CreatePublicNetworkCommand {

	return &CreatePublicNetworkCommand{receiver, publicNetworkCreate}
}
