package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// CreatePrivateNetworkWithQueryCommand represents command that configures new private network on the server with optional use of query parameter
type CreatePrivateNetworkWithQueryCommand struct {
	receiver             receiver.BMCSDK
	privateNetworkCreate networkapiclient.PrivateNetworkCreate
	query                dto.Query
}

// Execute runs CreatePrivateNetworkWithQueryCommand
func (command *CreatePrivateNetworkWithQueryCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	force := command.query.Force

	privateNetwork, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksAPI.PrivateNetworksPost(context.Background()).Force(force).PrivateNetworkCreate(command.privateNetworkCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return privateNetwork, nil
	}
	return nil, fmt.Errorf("CreatePrivateNetworkWithQueryCommand %s", errResolver.Error)
}

//NewCreatePrivateNetworkWithQueryCommand constructs new commmand of this type
func NewCreatePrivateNetworkWithQueryCommand(receiver receiver.BMCSDK, privateNetworkCreate networkapiclient.PrivateNetworkCreate, query dto.Query) *CreatePrivateNetworkWithQueryCommand {

	return &CreatePrivateNetworkWithQueryCommand{receiver, privateNetworkCreate, query}
}
