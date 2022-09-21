package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// CreatePrivateNetworkCommand represents command that configures new private network on the server
type CreatePrivateNetworkCommand struct {
	receiver             receiver.BMCSDK
	privateNetworkCreate networkapiclient.PrivateNetworkCreate
}

// Execute runs CreatePrivateNetworkCommand
func (command *CreatePrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(command.privateNetworkCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("CreatePrivateNetworkCommand %s", errResolver.Error)
}

//NewCreatePrivateNetworkCommand constructs new commmand of this type
func NewCreatePrivateNetworkCommand(receiver receiver.BMCSDK, privateNetworkCreate networkapiclient.PrivateNetworkCreate) *CreatePrivateNetworkCommand {

	return &CreatePrivateNetworkCommand{receiver, privateNetworkCreate}
}
