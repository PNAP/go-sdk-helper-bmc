package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// SetServerTagsCommand represents command that sets tags for specific server
type SetServerTagsCommand struct {
	receiver              receiver.BMCSDK
	serverID              string
	tagAssignmentRequests []bmcapiclient.TagAssignmentRequest
}

// Execute sets tags for specific server
func (command *SetServerTagsCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdTagsPut(context.Background(), command.serverID).TagAssignmentRequest(command.tagAssignmentRequests).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &server, nil
	}
	return nil, fmt.Errorf("SetServerTagsCommand %s", errResolver.Error)
}

//NewSetServerTagsCommand constructs new commmand of this type
func NewSetServerTagsCommand(receiver receiver.BMCSDK, serverID string, tagAssignmentRequests []bmcapiclient.TagAssignmentRequest) *SetServerTagsCommand {

	return &SetServerTagsCommand{receiver, serverID, tagAssignmentRequests}
}
