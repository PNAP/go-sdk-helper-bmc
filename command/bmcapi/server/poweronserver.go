package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// PowerOnServerCommand represents command that powers on specific server
type PowerOnServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute powers on specific server
func (command *PowerOnServerCommand) Execute() (*bmcapiclient.ActionResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdActionsPowerOnPost(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("PowerOnServerCommand %s", errResolver.Error)
}

//NewPowerOnServerCommand constructs new commmand of this type
func NewPowerOnServerCommand(receiver receiver.BMCSDK, serverID string) *PowerOnServerCommand {

	return &PowerOnServerCommand{receiver, serverID}
}
