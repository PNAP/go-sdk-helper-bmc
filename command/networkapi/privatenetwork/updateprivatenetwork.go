package privatenetwork

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	 networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// UpdatePrivateNetworkCommand represents command that updates private network on the server
type UpdatePrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	networkID string
	privateNetworkModify networkapiclient.PrivateNetworkModify
}


// Execute runs UpdatePrivateNetworkCommand
func (command *UpdatePrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksNetworkIdPut(context.Background(), command.networkID).PrivateNetworkModify(command.privateNetworkModify).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("UpdatePrivateNetworkCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

//NewUpdatePrivateNetworkCommand constructs new commmand of this type
func NewUpdatePrivateNetworkCommand(receiver receiver.BMCSDK, networkID string, privateNetworkModify networkapiclient.PrivateNetworkModify) *UpdatePrivateNetworkCommand {

	return &UpdatePrivateNetworkCommand{receiver, networkID, privateNetworkModify}
}