package sshkey

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// GetSshKeyCommand represents command that retrieves details about specific ssh key for the account
type GetSshKeyCommand struct {
	receiver receiver.BMCSDK
	sshKeyID string
}

// Execute runs GetSshKeyCommand
func (command *GetSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	sshKey, httpResponse, err := command.receiver.APIClient.SSHKeysAPI.SshKeysSshKeyIdGet(context.Background(), command.sshKeyID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return sshKey, nil
	}
	return nil, fmt.Errorf("GetSshKeyCommand %s", errResolver.Error)
}

//NewGetSshKeyCommand constructs new commmand of this type
func NewGetSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string) *GetSshKeyCommand {

	return &GetSshKeyCommand{receiver, sshKeyID}
}
