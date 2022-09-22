package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetServerCommand represents command that pulls details about specific server
type GetServerCommand struct {
	receiver receiver.BMCSDK
	serverID string
}

// Execute pulls details about specific server
func (command *GetServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdGet(context.Background(), command.serverID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("GetServerCommand %s", errResolver.Error)
}

// SetReceiver sets receiver to the command
func (command *GetServerCommand) SetReceiver(receiver receiver.BMCSDK) {
	command.receiver = receiver
}

// SetServerID sets server id to the command
func (command *GetServerCommand) SetServerID(id string) {
	command.serverID = id
}

//NewGetServerCommand constructs new commmand of this type
func NewGetServerCommand(receiver receiver.BMCSDK, serverID string) *GetServerCommand {

	return &GetServerCommand{receiver, serverID}
}
