package privatenetwork

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	 networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi"
)

// GetPrivateNetworksCommand represents command that lists all private networks on the account
type GetPrivateNetworksCommand struct {
	receiver receiver.BMCSDK
}


// Execute runs GetPrivateNetworksCommand
func (command *GetPrivateNetworksCommand) Execute() ([]networkapiclient.PrivateNetwork, error) {

	server, httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksGet(context.Background()).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("GetPrivateNetworksCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return server, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("GetPrivateNetworksCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewGetPrivateNetworksCommand constructs new commmand of this type
func NewGetPrivateNetworksCommand(receiver receiver.BMCSDK) *GetPrivateNetworksCommand {

	return &GetPrivateNetworksCommand{receiver}
}