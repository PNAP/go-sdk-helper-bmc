package privatenetwork

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	 networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// GetPrivateNetworkCommand represents command that retrieves details about speific private network on the account
type GetPrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	networkID string
}


// Execute runs GetPrivateNetworkCommand
func (command *GetPrivateNetworkCommand) Execute() (*networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksNetworkIdGet(context.Background(), command.networkID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("GetPrivateNetworkCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

//NewGetPrivateNetworkCommand constructs new commmand of this type
func NewGetPrivateNetworkCommand(receiver receiver.BMCSDK, networkID string) *GetPrivateNetworkCommand {

	return &GetPrivateNetworkCommand{receiver, networkID}
}