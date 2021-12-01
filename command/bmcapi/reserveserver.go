package bmcapi


import (
	"fmt"
	"github.com/PNAP/go-sdk-helper-bmc/dto"
	//"net/http"
	"context"
	 "github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// ReserveServerCommand represents command that reserves specific server
type ReserveServerCommand struct {
	receiver receiver.BMCSDK
	serverID  string
	serverReserve bmcapiclient.ServerReserve
}

// Execute reserve command on specific server
func (command *ReserveServerCommand) Execute() (*bmcapiclient.Server, error) {

	server, httpResponse, err := command.receiver.APIClient.ServersApi.ServersServerIdActionsReservePost(context.Background(), command.serverID).ServerReserve(command.serverReserve).Execute()

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

//NewReserveServerCommand constructs new commmand of this type
func NewReserveServerCommand(receiver receiver.BMCSDK, serverID string, serverReserve bmcapiclient.ServerReserve) *ReserveServerCommand {

	return &ReserveServerCommand{receiver, serverID, serverReserve}
}