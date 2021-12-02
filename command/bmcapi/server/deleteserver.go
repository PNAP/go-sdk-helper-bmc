package server

import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)


// DeleteServerCommand represents command that deletes specific server
type DeleteServerCommand struct {
	receiver receiver.BMCSDK
	serverID  string
}
// Execute deprovisions specific server
func (command *DeleteServerCommand) Execute() (*bmcapiclient.DeleteResult, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdDelete(context.Background(), command.serverID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("DeleteServerCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
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

//NewDeleteServerCommand constructs new commmand of this type
func NewDeleteServerCommand(receiver receiver.BMCSDK, serverID string) *DeleteServerCommand {

	return &DeleteServerCommand{receiver, serverID}
}