package sshkey

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// UpdateSshKeyCommand represents command that updates specific ssh key for the account
type UpdateSshKeyCommand struct {
	receiver     receiver.BMCSDK
	sshKeyID     string
	sshKeyUpdate bmcapiclient.SshKeyUpdate
}

// Execute runs UpdateSshKeyCommand
func (command *UpdateSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	sshKey, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysSshKeyIdPut(context.Background(), command.sshKeyID).SshKeyUpdate(command.sshKeyUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return sshKey, nil
	}
	return nil, fmt.Errorf("UpdateSshKeyCommand %s", errResolver.Error)
}

//NewUpdateSshKeyCommand constructs new commmand of this type
func NewUpdateSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string, sshKeyUpdate bmcapiclient.SshKeyUpdate) *UpdateSshKeyCommand {

	return &UpdateSshKeyCommand{receiver, sshKeyID, sshKeyUpdate}
}
