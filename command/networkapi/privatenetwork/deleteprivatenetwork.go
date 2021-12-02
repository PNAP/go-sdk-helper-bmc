package privatenetwork

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeletePrivateNetworkCommand represents command that deletes specific private network on the account
type DeletePrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	networkID string
}


// Execute runs DeletePrivateNetworkCommand
func (command *DeletePrivateNetworkCommand) Execute() (error) {

	httpResponse, err := command.receiver.NetworkAPIClient.PrivateNetworksApi.PrivateNetworksNetworkIdDelete(context.Background(), command.networkID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return err
		}
		return fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return error
		}
		return fmt.Errorf("DeletePrivateNetworkCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewDeletePrivateNetworkCommand constructs new commmand of this type
func NewDeletePrivateNetworkCommand(receiver receiver.BMCSDK, networkID string) *DeletePrivateNetworkCommand {

	return &DeletePrivateNetworkCommand{receiver, networkID}
}