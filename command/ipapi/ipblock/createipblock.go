package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// CreateIpBlockCommand represents command that adds a new IP Block to the account
type CreateIpBlockCommand struct {
	receiver      receiver.BMCSDK
	ipBlockCreate ipapiclient.IpBlockCreate
}

// Execute runs CreateIpBlockCommand
func (command *CreateIpBlockCommand) Execute() (*ipapiclient.IpBlock, error) {

	ipBlock, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksPost(context.Background()).IpBlockCreate(command.ipBlockCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &ipBlock, nil
	}
	return nil, fmt.Errorf("CreateIpBlockCommand %s", errResolver.Error)
}

//NewCreateIpBlockCommand constructs new commmand of this type
func NewCreateIpBlockCommand(receiver receiver.BMCSDK, ipBlockCreate ipapiclient.IpBlockCreate) *CreateIpBlockCommand {

	return &CreateIpBlockCommand{receiver, ipBlockCreate}
}
