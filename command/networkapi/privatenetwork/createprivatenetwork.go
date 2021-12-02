package privatenetwork

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	 networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// CreatePrivateNetworkCommand represents command that configures new private network on the server
type CreatePrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	privateNetworkCreate networkapiclient.PrivateNetworkCreate
}


// Execute runs CreatePrivateNetworkCommand
func (command *CreatePrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksPost(context.Background()).PrivateNetworkCreate(command.privateNetworkCreate).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("CreatePrivateNetworkCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

//NewCreatePrivateNetworkCommand constructs new commmand of this type
func NewCreatePrivateNetworkCommand(receiver receiver.BMCSDK, privateNetworkCreate networkapiclient.PrivateNetworkCreate) *CreatePrivateNetworkCommand {

	return &CreatePrivateNetworkCommand{receiver, privateNetworkCreate}
}