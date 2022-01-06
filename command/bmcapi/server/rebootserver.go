package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// RebootServerCommand represents command that reboots specific server
type RebootServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute reboots specific server
func (command *RebootServerCommand) Execute() (*bmcapiclient.ActionResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsRebootPost(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &result, nil
	}
	return nil, fmt.Errorf("RebootServerCommand %s", errResolver.Error)
}

//NewRebootServerCommand constructs new commmand of this type
func NewRebootServerCommand(receiver receiver.BMCSDK, serverID string) *RebootServerCommand {

	return &RebootServerCommand{receiver, serverID}
}
