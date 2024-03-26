package server

import (
	"fmt"
	"net/http"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// CreateServerCommand represents command that provisions new server
type CreateServerCommand struct {
	receiver     receiver.BMCSDK
	serverCreate bmcapiclient.ServerCreate
	query        *dto.Query
}

// Execute runs CreateServerCommand
func (command *CreateServerCommand) Execute() (*bmcapiclient.Server, error) {

	var server *bmcapiclient.Server
	var httpResponse *http.Response
	var err error

	if command.query != nil {

		force := command.query.Force

		server, httpResponse, err = command.receiver.APIClient.ServersAPI.ServersPost(context.Background()).Force(force).ServerCreate(command.serverCreate).Execute()
	} else {

		server, httpResponse, err = command.receiver.APIClient.ServersAPI.ServersPost(context.Background()).ServerCreate(command.serverCreate).Execute()
	}

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("CreateServerCommand %s", errResolver.Error)
}

//NewCreateServerCommand constructs new commmand of this type
func NewCreateServerCommand(receiver receiver.BMCSDK, serverCreate bmcapiclient.ServerCreate) *CreateServerCommand {

	return &CreateServerCommand{receiver, serverCreate, nil}
}

//NewCreateServerCommandWithQuery constructs new commmand of this type
func NewCreateServerCommandWithQuery(receiver receiver.BMCSDK, serverCreate bmcapiclient.ServerCreate, query *dto.Query) *CreateServerCommand {

	return &CreateServerCommand{receiver, serverCreate, query}
}
