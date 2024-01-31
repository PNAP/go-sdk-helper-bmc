package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// DeprovisionServerCommand represents command that deprovisions specific server
type DeprovisionServerCommand struct {
	receiver          receiver.BMCSDK
	serverID          string
	relinquishIpBlock bmcapiclient.RelinquishIpBlock
}

// Execute deprovisions specific server
func (command *DeprovisionServerCommand) Execute() (*string, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdActionsDeprovisionPost(context.Background(), command.serverID).RelinquishIpBlock(command.relinquishIpBlock).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &result, nil
	}
	return nil, fmt.Errorf("DeprovisionServerCommand %s", errResolver.Error)
}

//NewDeprovisionServerCommand constructs new commmand of this type
func NewDeprovisionServerCommand(receiver receiver.BMCSDK, serverID string, relinquishIpBlock bmcapiclient.RelinquishIpBlock) *DeprovisionServerCommand {

	return &DeprovisionServerCommand{receiver, serverID, relinquishIpBlock}
}
