package publicnetwork

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// RemoveIpBlockFromPublicNetworkWithQueryCommand represents command that removes an IP Block from specific public network with optional use of query parameter
type RemoveIpBlockFromPublicNetworkWithQueryCommand struct {
	receiver  receiver.BMCSDK
	networkID string
	ipBlockID string
	query     dto.Query
}

// Execute runs RemoveIpBlockFromPublicNetworkWithQueryCommand
func (command *RemoveIpBlockFromPublicNetworkWithQueryCommand) Execute() (*string, error) {

	force := command.query.Force

	response, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksIpBlockIdDelete(context.Background(), command.networkID, command.ipBlockID).Force(force).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &response, nil
	}
	return nil, fmt.Errorf("RemoveIpBlockFromPublicNetworkWithQueryCommand %s", errResolver.Error)
}

//NewRemoveIpBlockFromPublicNetworkWithQueryCommand constructs new commmand of this type
func NewRemoveIpBlockFromPublicNetworkWithQueryCommand(receiver receiver.BMCSDK, networkID string, ipBlockID string, query dto.Query) *RemoveIpBlockFromPublicNetworkWithQueryCommand {

	return &RemoveIpBlockFromPublicNetworkWithQueryCommand{receiver, networkID, ipBlockID, query}
}
