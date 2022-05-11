package publicnetwork

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// RemoveIpBlockFromPublicNetworkCommand represents command that removes an IP Block from specific public network
type RemoveIpBlockFromPublicNetworkCommand struct {
	receiver  receiver.BMCSDK
	networkID string
	ipBlockID string
}

// Execute runs RemoveIpBlockFromPublicNetworkCommand
func (command *RemoveIpBlockFromPublicNetworkCommand) Execute() (*string, error) {

	response, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksApi.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(context.Background(), command.networkID, command.ipBlockID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &response, nil
	}
	return nil, fmt.Errorf("RemoveIpBlockFromPublicNetworkCommand %s", errResolver.Error)
}

//NewRemoveIpBlockFromPublicNetworkCommand constructs new commmand of this type
func NewRemoveIpBlockFromPublicNetworkCommand(receiver receiver.BMCSDK, networkID string, ipBlockID string) *RemoveIpBlockFromPublicNetworkCommand {

	return &RemoveIpBlockFromPublicNetworkCommand{receiver, networkID, ipBlockID}
}
