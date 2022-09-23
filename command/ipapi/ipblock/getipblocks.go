package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
)

// GetIpBlocksCommand represents command that retrieves all IP Blocks beloging to the account
type GetIpBlocksCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetIpBlocksCommand
func (command *GetIpBlocksCommand) Execute() ([]ipapiclient.IpBlock, error) {

	tags, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return tags, nil
	}
	return nil, fmt.Errorf("GetIpBlocksCommand %s", errResolver.Error)
}

//NewGetIpBlocksCommand constructs new commmand of this type
func NewGetIpBlocksCommand(receiver receiver.BMCSDK) *GetIpBlocksCommand {

	return &GetIpBlocksCommand{receiver}
}
