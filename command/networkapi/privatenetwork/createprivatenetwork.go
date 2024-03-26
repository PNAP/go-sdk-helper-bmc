package privatenetwork

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// CreatePrivateNetworkCommand represents command that configures new private network on the server
type CreatePrivateNetworkCommand struct {
	receiver             receiver.BMCSDK
	privateNetworkCreate networkapiclient.PrivateNetworkCreate
	query                *dto.Query
}

// Execute runs CreatePrivateNetworkCommand
func (command *CreatePrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	var privateNetwork *networkapiclient.PrivateNetwork
	var httpResponse *http.Response
	var err error

	if command.query != nil {

		force := command.query.Force

		privateNetwork, httpResponse, err = command.receiver.NetworkAPIClient.PrivateNetworksAPI.PrivateNetworksPost(context.Background()).Force(force).
			PrivateNetworkCreate(command.privateNetworkCreate).Execute()
	} else {

		privateNetwork, httpResponse, err = command.receiver.NetworkAPIClient.PrivateNetworksAPI.PrivateNetworksPost(context.Background()).
			PrivateNetworkCreate(command.privateNetworkCreate).Execute()
	}

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return privateNetwork, nil
	}
	return nil, fmt.Errorf("CreatePrivateNetworkCommand %s", errResolver.Error)
}

//NewCreatePrivateNetworkCommand constructs new commmand of this type
func NewCreatePrivateNetworkCommand(receiver receiver.BMCSDK, privateNetworkCreate networkapiclient.PrivateNetworkCreate) *CreatePrivateNetworkCommand {

	return &CreatePrivateNetworkCommand{receiver, privateNetworkCreate, nil}
}

//NewCreatePrivateNetworkCommandWithQuery constructs new commmand of this type
func NewCreatePrivateNetworkCommandWithQuery(receiver receiver.BMCSDK, privateNetworkCreate networkapiclient.PrivateNetworkCreate, query *dto.Query) *CreatePrivateNetworkCommand {

	return &CreatePrivateNetworkCommand{receiver, privateNetworkCreate, query}
}
