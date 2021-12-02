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
	receiver receiver.BMCSDK
	sshKeyCreate bmcapiclient.SshKeyCreate
}


// Execute runs CreateSshKeyCommand
func (command *CreateSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	server, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysPost(context.Background()).SshKeyCreate(command.sshKeyCreate).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("CreateSshKeyCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return &server, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewCreateSshKeyCommand constructs new commmand of this type
func NewCreateSshKeyCommand(receiver receiver.BMCSDK, sshKeyCreate bmcapiclient.SshKeyCreate) *CreateSshKeyCommand {

	return &CreateSshKeyCommand{receiver, sshKeyCreate}
}