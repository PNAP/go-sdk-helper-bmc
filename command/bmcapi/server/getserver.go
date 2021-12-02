package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetServerCommand represents command that pulls details about specific server
type GetServerCommand struct {
	receiver receiver.BMCSDK
	serverID  string
}

// Execute pulls details about specific server
func (command *GetServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdGet(context.Background(), command.serverID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("GetServerCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

 // SetReceiver sets receiver to the command
func (command *GetServerCommand) SetReceiver(receiver receiver.BMCSDK) {
	command.receiver = receiver
}

// SetServerID sets server id to the command
func (command *GetServerCommand) SetServerID(id string) {
	command.serverID = id
} 

//NewGetServerCommand constructs new commmand of this type
func NewGetServerCommand(receiver receiver.BMCSDK, serverID string) *GetServerCommand {

	return &GetServerCommand{receiver, serverID}
}
