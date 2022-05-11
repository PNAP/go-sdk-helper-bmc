package publicnetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeletePublicNetworkCommand represents command that deletes specific public network on the account
type DeletePublicNetworkCommand struct {
	receiver  receiver.BMCSDK
	networkID string
}

// Execute runs DeletePublicNetworkCommand
func (command *DeletePublicNetworkCommand) Execute() error {

	httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksApi.PublicNetworksNetworkIdDelete(context.Background(), command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return nil
	}
	return fmt.Errorf("DeletePublicNetworkCommand %s", errResolver.Error)
}

//NewDeletePublicNetworkCommand constructs new commmand of this type
func NewDeletePublicNetworkCommand(receiver receiver.BMCSDK, networkID string) *DeletePublicNetworkCommand {

	return &DeletePublicNetworkCommand{receiver, networkID}
}
