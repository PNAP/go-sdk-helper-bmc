package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// ResetServerCommand represents command that resets specific server
type ResetServerCommand struct {
	receiver receiver.BMCSDK
	serverID  string
	serverReset bmcapiclient.ServerReset
}

// Execute resets specific server
func (command *ResetServerCommand) Execute() (*bmcapiclient.ResetResult, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsResetPost(context.Background(), command.serverID).ServerReset(command.serverReset).Execute()

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

//NewResetServerCommand constructs new commmand of this type
func NewResetServerCommand(receiver receiver.BMCSDK, serverID string, serverReset bmcapiclient.ServerReset) *ResetServerCommand {

	return &ResetServerCommand{receiver, serverID, serverReset}
}