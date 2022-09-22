package tag

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

// UpdateTagCommand represents command that updates a specific tag for the account
type UpdateTagCommand struct {
	receiver  receiver.BMCSDK
	tagID     string
	tagUpdate tagapiclient.TagUpdate
}

// Execute runs UpdateTagCommand
func (command *UpdateTagCommand) Execute() (*tagapiclient.Tag, error) {

	tag, httpResponse, err := command.receiver.TagAPIClient.TagsApi.TagsTagIdPatch(context.Background(), command.tagID).TagUpdate(command.tagUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return tag, nil
	}
	return nil, fmt.Errorf("UpdateTagCommand %s", errResolver.Error)
}

//NewUpdateTagCommand constructs new commmand of this type
func NewUpdateTagCommand(receiver receiver.BMCSDK, tagID string, tagUpdate tagapiclient.TagUpdate) *UpdateTagCommand {

	return &UpdateTagCommand{receiver, tagID, tagUpdate}
}
