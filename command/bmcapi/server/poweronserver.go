package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// PowerOnServerCommand represents command that powers on specific server
type PowerOnServerCommand struct {
	receiver receiver.BMCSDK
	serverID  string
}

// Execute powers on specific server
func (command *PowerOnServerCommand) Execute() (*bmcapiclient.ActionResult, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsPowerOnPost(context.Background(), command.serverID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("PowerOnServerCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

//NewPowerOnServerCommand constructs new commmand of this type
func NewPowerOnServerCommand(receiver receiver.BMCSDK, serverID string) *PowerOnServerCommand {

	return &PowerOnServerCommand{receiver, serverID}
}