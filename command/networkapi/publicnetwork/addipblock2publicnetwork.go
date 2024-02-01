package publicnetwork

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// AddIpBlock2PublicNetworkCommand represents command that adds an IP Block to specific public network
type AddIpBlock2PublicNetworkCommand struct {
	receiver             receiver.BMCSDK
	networkID            string
	publicNetworkIpBlock networkapiclient.PublicNetworkIpBlock
}

// Execute runs AddIpBlock2PublicNetworkCommand
func (command *AddIpBlock2PublicNetworkCommand) Execute() (*networkapiclient.PublicNetworkIpBlock, error) {

	ipBlock, httpResponse, err := command.receiver.NetworkAPIClient.PublicNetworksAPI.PublicNetworksNetworkIdIpBlocksPost(context.Background(), command.networkID).PublicNetworkIpBlock(command.publicNetworkIpBlock).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return ipBlock, nil
	}
	return nil, fmt.Errorf("AddIpBlock2PublicNetworkCommand %s", errResolver.Error)
}

//NewAddIpBlock2PublicNetworkCommand constructs new commmand of this type
func NewAddIpBlock2PublicNetworkCommand(receiver receiver.BMCSDK, networkID string, publicNetworkIpBlock networkapiclient.PublicNetworkIpBlock) *AddIpBlock2PublicNetworkCommand {

	return &AddIpBlock2PublicNetworkCommand{receiver, networkID, publicNetworkIpBlock}
}
