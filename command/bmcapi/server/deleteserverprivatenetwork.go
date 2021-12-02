package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeleteServerPrivateNetworkCommand represents command that removes private network from networking devices for specific server
type DeleteServerPrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	serverID  string
	networkID  string
}

// Execute runs DeleteServerPrivateNetwork command for specific server
func (command *DeleteServerPrivateNetworkCommand) Execute() (*string, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.DeletePrivateNetwork(context.Background(), command.serverID, command.networkID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return &server, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("DeleteServerPrivateNetworkCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewDeleteServerPrivateNetworkCommand constructs new commmand of this type
func NewDeleteServerPrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, networkID  string) *DeleteServerPrivateNetworkCommand {

	return &DeleteServerPrivateNetworkCommand{receiver, serverID, networkID}
}