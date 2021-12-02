package server


import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// CreateServerCommand represents command that provisions new server
type CreateServerCommand struct {
	receiver receiver.BMCSDK
	serverCreate bmcapiclient.ServerCreate
}


// Execute runs CreateServerCommand 
func (command *CreateServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersPost(context.Background()).ServerCreate(command.serverCreate).Execute()

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

//NewCreateServerCommand constructs new commmand of this type
func NewCreateServerCommand(receiver receiver.BMCSDK, serverCreate bmcapiclient.ServerCreate) *CreateServerCommand {

	return &CreateServerCommand{receiver, serverCreate}
}