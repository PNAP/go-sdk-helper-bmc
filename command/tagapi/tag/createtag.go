package tag

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
)

// CreateTagCommand represents command that adds a new tag to the account
type CreateTagCommand struct {
	receiver  receiver.BMCSDK
	tagCreate tagapiclient.TagCreate
}

// Execute runs CreateTagCommand
func (command *CreateTagCommand) Execute() (*tagapiclient.Tag, error) {

	tag, httpResponse, err := command.receiver.TagAPIClient.TagsAPI.TagsPost(context.Background()).TagCreate(command.tagCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return tag, nil
	}
	return nil, fmt.Errorf("CreateTagCommand %s", errResolver.Error)
}

//NewCreateTagCommand constructs new commmand of this type
func NewCreateTagCommand(receiver receiver.BMCSDK, tagCreate tagapiclient.TagCreate) *CreateTagCommand {

	return &CreateTagCommand{receiver, tagCreate}
}
