package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// ShutDownServerCommand represents command that shuts down specific server
type ShutDownServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute shuts down specific server
func (command *ShutDownServerCommand) Execute() (*bmcapiclient.ActionResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsShutdownPost(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("ShutDownServerCommand %s", errResolver.Error)
}

//NewShutDownServerCommand constructs new commmand of this type
func NewShutDownServerCommand(receiver receiver.BMCSDK, serverID string) *ShutDownServerCommand {

	return &ShutDownServerCommand{receiver, serverID}
}
