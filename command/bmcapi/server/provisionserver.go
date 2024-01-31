package server

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// ProvisionServerCommand represents command that provisions a reserved server
type ProvisionServerCommand struct {
	receiver        receiver.BMCSDK
	serverID        string
	serverProvision bmcapiclient.ServerProvision
	query           dto.Query
}

// Execute runs ProvisionServerCommand
func (command *ProvisionServerCommand) Execute() (*bmcapiclient.Server, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdActionsProvisionPost(context.Background(), command.serverID).Force(force).
		ServerProvision(command.serverProvision).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("ProvisionServerCommand %s", errResolver.Error)
}

//NewProvisionServerCommand constructs new commmand of this type
func NewProvisionServerCommand(receiver receiver.BMCSDK, serverID string, serverProvision bmcapiclient.ServerProvision, query dto.Query) *ProvisionServerCommand {

	return &ProvisionServerCommand{receiver, serverID, serverProvision, query}
}
