package tag

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

// GetTagsCommand represents command that retrieves all tags beloging to the account
type GetTagsCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetTagsCommand
func (command *GetTagsCommand) Execute() ([]tagapiclient.Tag, error) {

	tags, httpResponse, err := command.receiver.TagAPIClient.TagsApi.TagsGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return tags, nil
	}
	return nil, fmt.Errorf("GetTagsCommand %s", errResolver.Error)
}

//NewGetTagsCommand constructs new commmand of this type
func NewGetTagsCommand(receiver receiver.BMCSDK) *GetTagsCommand {

	return &GetTagsCommand{receiver}
}
