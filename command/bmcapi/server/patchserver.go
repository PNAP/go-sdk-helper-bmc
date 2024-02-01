package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// PatchServerCommand represents command that updates specific server
type PatchServerCommand struct {
	receiver    receiver.BMCSDK
	serverID    string
	serverPatch bmcapiclient.ServerPatch
}

// Execute updates specific server
func (command *PatchServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdPatch(context.Background(), command.serverID).ServerPatch(command.serverPatch).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("PatchServerCommand %s", errResolver.Error)
}

//NewPatchServerCommand constructs new commmand of this type
func NewPatchServerCommand(receiver receiver.BMCSDK, serverID string, serverPatch bmcapiclient.ServerPatch) *PatchServerCommand {

	return &PatchServerCommand{receiver, serverID, serverPatch}
}
