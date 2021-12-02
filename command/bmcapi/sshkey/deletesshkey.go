package sshkey

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// DeleteSshKeyCommand represents command that deletes specific ssh key from the account
type DeleteSshKeyCommand struct {
	receiver receiver.BMCSDK
	sshKeyID string
}


// Execute runs DeleteSshKeyCommand
func (command *DeleteSshKeyCommand) Execute() (*bmcapiclient.DeleteSshKeyResult, error) {

	result, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysSshKeyIdDelete(context.Background(), command.sshKeyID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return &result, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewDeleteSshKeyCommand constructs new commmand of this type
func NewDeleteSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string) *DeleteSshKeyCommand {

	return &DeleteSshKeyCommand{receiver, sshKeyID}
}