package sshkey

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// DeleteSshKeyCommand represents command that deletes specific ssh key from the account
type DeleteSshKeyCommand struct {
	receiver receiver.BMCSDK
	sshKeyID string
}

// Execute runs DeleteSshKeyCommand
func (command *DeleteSshKeyCommand) Execute() (*bmcapiclient.DeleteSshKeyResult, error) {

	result, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysSshKeyIdDelete(context.Background(), command.sshKeyID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("DeleteSshKeyCommand %s", errResolver.Error)
}

//NewDeleteSshKeyCommand constructs new commmand of this type
func NewDeleteSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string) *DeleteSshKeyCommand {

	return &DeleteSshKeyCommand{receiver, sshKeyID}
}
