package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// RemoveIpBlockCommand represents command that removes IP Block from specific server
type RemoveIpBlockCommand struct {
	receiver          receiver.BMCSDK
	serverID          string
	ipBlockID         string
	relinquishIpBlock bmcapiclient.RelinquishIpBlock
}

// Execute updates specific server
func (command *RemoveIpBlockCommand) Execute() (*string, error) {

	response, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdIpBlocksIpBlockIdDelete(context.Background(), command.serverID, command.ipBlockID).RelinquishIpBlock(command.relinquishIpBlock).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &response, nil
	}
	return nil, fmt.Errorf("RemoveIpBlockCommand %s", errResolver.Error)
}

//NewRemoveIpBlockCommand constructs new commmand of this type
func NewRemoveIpBlockCommand(receiver receiver.BMCSDK, serverID string, ipBlockID string, relinquishIpBlock bmcapiclient.RelinquishIpBlock) *RemoveIpBlockCommand {

	return &RemoveIpBlockCommand{receiver, serverID, ipBlockID, relinquishIpBlock}
}
