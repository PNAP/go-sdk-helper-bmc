package bmcapi


import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetServersCommand represents command that pulls details about all servers
type GetServersCommand struct {
	receiver receiver.BMCSDK
}

// Execute pulls details about specific server
func (command *GetServersCommand) Execute() ([]bmcapiclient.Server, error) {

	servers, httpResponse, err := command.receiver.APIClient.ServersApi.ServersGet(context.Background()).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, err
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	
		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300{
		return servers, nil
	} else{
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil{
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}
	
}

//NewGetServersCommand constructs new commmand of this type
func NewGetServersCommand(receiver receiver.BMCSDK) *GetServersCommand {

	return &GetServersCommand{receiver}
}
