package sshkey

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// GetSshKeysCommand represents command that lists all ssh keys for the account
type GetSshKeysCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetSshKeysCommand
func (command *GetSshKeysCommand) Execute() ([]bmcapiclient.SshKey, error) {

	sshKeys, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return sshKeys, nil
	}
	return nil, fmt.Errorf("GetSshKeysCommand %s", errResolver.Error)
}

//NewGetSshKeysCommand constructs new commmand of this type
func NewGetSshKeysCommand(receiver receiver.BMCSDK) *GetSshKeysCommand {

	return &GetSshKeysCommand{receiver}
}
