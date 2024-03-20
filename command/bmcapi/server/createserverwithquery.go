package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// CreateServerWithQueryCommand represents command that provisions new server with optional use of query parameter
type CreateServerWithQueryCommand struct {
	receiver     receiver.BMCSDK
	serverCreate bmcapiclient.ServerCreate
	query        dto.Query
}

// Execute runs CreateServerWithQueryCommand
func (command *CreateServerWithQueryCommand) Execute() (*bmcapiclient.Server, error) {

	force := command.query.Force

	server, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersPost(context.Background()).Force(force).ServerCreate(command.serverCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("CreateServerWithQueryCommand %s", errResolver.Error)
}

//NewCreateServerWithQueryCommand constructs new commmand of this type
func NewCreateServerWithQueryCommand(receiver receiver.BMCSDK, serverCreate bmcapiclient.ServerCreate, query dto.Query) *CreateServerWithQueryCommand {

	return &CreateServerWithQueryCommand{receiver, serverCreate, query}
}
