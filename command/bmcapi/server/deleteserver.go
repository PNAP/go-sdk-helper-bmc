package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// DeleteServerCommand represents command that deletes specific server
type DeleteServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute deprovisions specific server
func (command *DeleteServerCommand) Execute() (*bmcapiclient.DeleteResult, error) {

	result, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdDelete(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("DeleteServerCommand %s", errResolver.Error)
}

//NewDeleteServerCommand constructs new commmand of this type
func NewDeleteServerCommand(receiver receiver.BMCSDK, serverID string) *DeleteServerCommand {

	return &DeleteServerCommand{receiver, serverID}
}
