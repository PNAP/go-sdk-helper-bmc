package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// GetIpBlockCommand represents command that retrieves a specific IP Block belonging to the account
type GetIpBlockCommand struct {
	receiver  receiver.BMCSDK
	ipBlockID string
}

// Execute runs GetIpBlockCommand
func (command *GetIpBlockCommand) Execute() (*ipapiclient.IpBlock, error) {

	tag, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksIpBlockIdGet(context.Background(), command.ipBlockID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &tag, nil
	}
	return nil, fmt.Errorf("GetIpBlockCommand %s", errResolver.Error)
}

//NewGetIpBlockCommand constructs new commmand of this type
func NewGetIpBlockCommand(receiver receiver.BMCSDK, ipBlockID string) *GetIpBlockCommand {

	return &GetIpBlockCommand{receiver, ipBlockID}
}
