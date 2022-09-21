package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// ResetServerCommand represents command that resets specific server
type ResetServerCommand struct {
	receiver    receiver.BMCSDK
	serverID    string
	serverReset bmcapiclient.ServerReset
}

// Execute resets specific server
func (command *ResetServerCommand) Execute() (*bmcapiclient.ResetResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsResetPost(context.Background(), command.serverID).ServerReset(command.serverReset).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("ResetServerCommand %s", errResolver.Error)
}

//NewResetServerCommand constructs new commmand of this type
func NewResetServerCommand(receiver receiver.BMCSDK, serverID string, serverReset bmcapiclient.ServerReset) *ResetServerCommand {

	return &ResetServerCommand{receiver, serverID, serverReset}
}
