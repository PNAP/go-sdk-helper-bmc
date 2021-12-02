package sshkey

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetSshKeysCommand represents command that lists all ssh keys for the account
type GetSshKeysCommand struct {
	receiver receiver.BMCSDK
	
}


// Execute runs GetSshKeysCommand
func (command *GetSshKeysCommand) Execute() ([]bmcapiclient.SshKey, error) {

	server, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysGet(context.Background()).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return server, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewGetSshKeysCommand constructs new commmand of this type
func NewGetSshKeysCommand(receiver receiver.BMCSDK) *GetSshKeysCommand {

	return &GetSshKeysCommand{receiver}
}