package server

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetServersCommand represents command that pulls details about all servers
type GetServersCommand struct {
	receiver receiver.BMCSDK
}

// Execute pulls details about all servers
func (command *GetServersCommand) Execute() ([]bmcapiclient.Server, error) {

	servers, httpResponse, err := command.receiver.APIClient.ServersApi.ServersGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return servers, nil
	}
	return nil, fmt.Errorf("GetServersCommand %s", errResolver.Error)
}

//NewGetServersCommand constructs new commmand of this type
func NewGetServersCommand(receiver receiver.BMCSDK) *GetServersCommand {

	return &GetServersCommand{receiver}
}
