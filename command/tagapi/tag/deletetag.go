package tag

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	tagapiclient "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

// DeleteTagCommand represents command that deletes a specific tag belonging to the account
type DeleteTagCommand struct {
	receiver receiver.BMCSDK
	tagID    string
}

// Execute runs DeleteTagCommand
func (command *DeleteTagCommand) Execute() (*tagapiclient.DeleteResult, error) {

	result, httpResponse, err := command.receiver.TagAPIClient.TagsApi.TagsTagIdDelete(context.Background(), command.tagID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &result, nil
	}
	return nil, fmt.Errorf("DeleteTagCommand %s", errResolver.Error)
}

//NewDeleteTagCommand constructs new commmand of this type
func NewDeleteTagCommand(receiver receiver.BMCSDK, tagID string) *DeleteTagCommand {

	return &DeleteTagCommand{receiver, tagID}
}
