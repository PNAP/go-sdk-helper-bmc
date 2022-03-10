package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// AddIpBlockCommand represents command that adds an IP Block to specific server
type AddIpBlockCommand struct {
	receiver      receiver.BMCSDK
	serverID      string
	serverIpBlock bmcapiclient.ServerIpBlock
}

// Execute adds an IP Block to specific server
func (command *AddIpBlockCommand) Execute() (*bmcapiclient.ServerIpBlock, error) {

	ipBlock, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdIpBlocksPost(context.Background(), command.serverID).ServerIpBlock(command.serverIpBlock).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &ipBlock, nil
	}
	return nil, fmt.Errorf("AddIpBlockCommand %s", errResolver.Error)
}

//NewAddIpBlockCommand constructs new commmand of this type
func NewAddIpBlockCommand(receiver receiver.BMCSDK, serverID string, serverIpBlock bmcapiclient.ServerIpBlock) *AddIpBlockCommand {

	return &AddIpBlockCommand{receiver, serverID, serverIpBlock}
}
