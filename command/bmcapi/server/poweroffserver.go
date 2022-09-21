package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// PowerOffServerCommand represents command that powers off specific server
type PowerOffServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute powers off specific server
func (command *PowerOffServerCommand) Execute() (*bmcapiclient.ActionResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsPowerOffPost(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("PowerOffServerCommand %s", errResolver.Error)
}

//NewPowerOffServerCommand constructs new commmand of this type
func NewPowerOffServerCommand(receiver receiver.BMCSDK, serverID string) *PowerOffServerCommand {

	return &PowerOffServerCommand{receiver, serverID}
}
