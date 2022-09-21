package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// DeleteIpBlockCommand represents command that deletes a specific IP Block belonging to the account
type DeleteIpBlockCommand struct {
	receiver  receiver.BMCSDK
	ipBlockID string
}

// Execute runs DeleteIpBlockCommand
func (command *DeleteIpBlockCommand) Execute() (*ipapiclient.DeleteIpBlockResult, error) {

	result, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksIpBlockIdDelete(context.Background(), command.ipBlockID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("DeleteIpBlockCommand %s", errResolver.Error)
}

//NewDeleteIpBlockCommand constructs new commmand of this type
func NewDeleteIpBlockCommand(receiver receiver.BMCSDK, ipBlockID string) *DeleteIpBlockCommand {

	return &DeleteIpBlockCommand{receiver, ipBlockID}
}
