package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// PatchIpBlockCommand represents command that updates the details of a specific IP Block belonging to the account
type PatchIpBlockCommand struct {
	receiver     receiver.BMCSDK
	ipBlockID    string
	ipBlockPatch ipapiclient.IpBlockPatch
}

// Execute runs PatchIpBlockCommand
func (command *PatchIpBlockCommand) Execute() (*ipapiclient.IpBlock, error) {

	ipBlock, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksIpBlockIdPatch(context.Background(), command.ipBlockID).IpBlockPatch(command.ipBlockPatch).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return ipBlock, nil
	}
	return nil, fmt.Errorf("PatchIpBlockCommand %s", errResolver.Error)
}

//NewPatchIpBlockCommand constructs new commmand of this type
func NewPatchIpBlockCommand(receiver receiver.BMCSDK, ipBlockID string, ipBlockPatch ipapiclient.IpBlockPatch) *PatchIpBlockCommand {

	return &PatchIpBlockCommand{receiver, ipBlockID, ipBlockPatch}
}
