package ipblock

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	ipapiclient "github.com/phoenixnap/go-sdk-bmc/ipapi"
)

// PutTagsIpBlockCommand represents command that overwrites tags assigned for specific IP Block belonging to the account
// and unassigns any tags not part of the request.
type PutTagsIpBlockCommand struct {
	receiver             receiver.BMCSDK
	ipBlockID            string
	tagAssignmentRequest []ipapiclient.TagAssignmentRequest
}

// Execute runs PutTagsIpBlockCommand
func (command *PutTagsIpBlockCommand) Execute() (*ipapiclient.IpBlock, error) {

	ipBlock, httpResponse, err := command.receiver.IpBlockAPIClient.IPBlocksApi.IpBlocksIpBlockIdTagsPut(context.Background(), command.ipBlockID).TagAssignmentRequest(command.tagAssignmentRequest).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &ipBlock, nil
	}
	return nil, fmt.Errorf("PutTagsIpBlockCommand %s", errResolver.Error)
}

//NewPutTagsIpBlockCommand constructs new commmand of this type
func NewPutTagsIpBlockCommand(receiver receiver.BMCSDK, ipBlockID string, tagAssignmentRequest []ipapiclient.TagAssignmentRequest) *PutTagsIpBlockCommand {

	return &PutTagsIpBlockCommand{receiver, ipBlockID, tagAssignmentRequest}
}
