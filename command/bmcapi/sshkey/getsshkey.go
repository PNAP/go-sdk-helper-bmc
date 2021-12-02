package sshkey




import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetSshKeyCommand represents command that retrieves details about specific ssh key for the account
type GetSshKeyCommand struct {
	receiver receiver.BMCSDK
	sshKeyID string
}


// Execute runs GetSshKeyCommand
func (command *GetSshKeyCommand) Execute() (*bmcapiclient.SshKey, error) {

	sshKey, httpResponse, err := command.receiver.APIClient.SSHKeysApi.SshKeysSshKeyIdGet(context.Background(), command.sshKeyID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("GetSshKeyCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return &sshKey, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewGetSshKeyCommand constructs new commmand of this type
func NewGetSshKeyCommand(receiver receiver.BMCSDK, sshKeyID string) *GetSshKeyCommand {

	return &GetSshKeyCommand{receiver, sshKeyID}
}