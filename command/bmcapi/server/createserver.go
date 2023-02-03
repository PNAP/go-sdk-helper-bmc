package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// CreateServerCommand represents command that provisions new server
type CreateServerCommand struct {
	receiver          receiver.BMCSDK
	serverCreate      bmcapiclient.ServerCreate
	createServerQuery dto.CreateServerQuery
}

// Execute runs CreateServerCommand
func (command *CreateServerCommand) Execute() (*bmcapiclient.Server, error) {

	force := command.createServerQuery.Force

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersPost(context.Background()).Force(force).ServerCreate(command.serverCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return server, nil
	}
	return nil, fmt.Errorf("CreateServerCommand %s", errResolver.Error)
}

//NewCreateServerCommand constructs new commmand of this type
func NewCreateServerCommand(receiver receiver.BMCSDK, serverCreate bmcapiclient.ServerCreate, createServerQuery dto.CreateServerQuery) *CreateServerCommand {

	return &CreateServerCommand{receiver, serverCreate, createServerQuery}
}
