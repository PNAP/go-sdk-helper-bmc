package tag

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
)

// GetTagCommand represents command that retrieves a specific tag belonging to the account
type GetTagCommand struct {
	receiver receiver.BMCSDK
	tagID    string
}

// Execute runs GetTagCommand
func (command *GetTagCommand) Execute() (*tagapiclient.Tag, error) {

	tag, httpResponse, err := command.receiver.TagAPIClient.TagsApi.TagsTagIdGet(context.Background(), command.tagID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return tag, nil
	}
	return nil, fmt.Errorf("GetTagCommand %s", errResolver.Error)
}

//NewGetTagCommand constructs new commmand of this type
func NewGetTagCommand(receiver receiver.BMCSDK, tagID string) *GetTagCommand {

	return &GetTagCommand{receiver, tagID}
}
