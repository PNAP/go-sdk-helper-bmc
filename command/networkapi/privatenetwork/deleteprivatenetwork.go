package privatenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeletePrivateNetworkCommand represents command that deletes specific private network on the account
type DeletePrivateNetworkCommand struct {
	receiver  receiver.BMCSDK
	networkID string
}

// Execute runs DeletePrivateNetworkCommand
func (command *DeletePrivateNetworkCommand) Execute() error {

	httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksAPI.PrivateNetworksNetworkIdDelete(context.Background(), command.networkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return nil
	}
	return fmt.Errorf("DeletePrivateNetworkCommand %s", errResolver.Error)
}

//NewDeletePrivateNetworkCommand constructs new commmand of this type
func NewDeletePrivateNetworkCommand(receiver receiver.BMCSDK, networkID string) *DeletePrivateNetworkCommand {

	return &DeletePrivateNetworkCommand{receiver, networkID}
}
