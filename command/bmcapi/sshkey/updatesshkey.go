package sshkey

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// UpdateSshKeyCommand represents command that updates specific ssh key for the account
type UpdateSshKeyCommand struct {
	receiver receiver.BMCSDK
	sshKeyID string
	sshKeyUpdate bmcapiclient.SshKeyUpdate
}


// Execute runs UpdateSshKeyCommand
func (command *UpdateSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	key, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysSshKeyIdPut(context.Background(), command.sshKeyID).SshKeyUpdate(command.sshKeyUpdate).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return &key, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewUpdateSshKeyCommand constructs new commmand of this type
func NewUpdateSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string, sshKeyUpdate bmcapiclient.SshKeyUpdate) *UpdateSshKeyCommand {

	return &UpdateSshKeyCommand{receiver, sshKeyID, sshKeyUpdate}
}