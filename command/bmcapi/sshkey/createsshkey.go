package sshkey

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// CreateSshKeyCommand represents command that adds new ssh key to the account
type CreateSshKeyCommand struct {
	receiver     receiver.BMCSDK
	sshKeyCreate bmcapiclient.SshKeyCreate
}

// Execute runs CreateSshKeyCommand
func (command *CreateSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	sshKey, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysPost(context.Background()).SshKeyCreate(command.sshKeyCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return sshKey, nil
	}
	return nil, fmt.Errorf("CreateSshKeyCommand %s", errResolver.Error)
}

//NewCreateSshKeyCommand constructs new commmand of this type
func NewCreateSshKeyCommand(receiver receiver.BMCSDK, sshKeyCreate bmcapiclient.SshKeyCreate) *CreateSshKeyCommand {

	return &CreateSshKeyCommand{receiver, sshKeyCreate}
}
