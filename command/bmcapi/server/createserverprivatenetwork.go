package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// CreateServerPrivateNetworkCommand represents command that configures private network on specific server
type CreateServerPrivateNetworkCommand struct {
	receiver receiver.BMCSDK
	serverID  string
	serverPrivateNetwork  bmcapiclient.ServerPrivateNetwork
}

// Execute con command on specific server
func (command *CreateServerPrivateNetworkCommand) Execute() (*bmcapiclient.ServerPrivateNetwork, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdPrivateNetworksPost(context.Background(), command.serverID).ServerPrivateNetwork(command.serverPrivateNetwork).Execute()

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
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewCreateServerPrivateNetworkCommand constructs new commmand of this type
func NewCreateServerPrivateNetworkCommand(receiver receiver.BMCSDK, serverID string, serverPrivateNetwork bmcapiclient.ServerPrivateNetwork) *CreateServerPrivateNetworkCommand {

	return &CreateServerPrivateNetworkCommand{receiver, serverID, serverPrivateNetwork}
}