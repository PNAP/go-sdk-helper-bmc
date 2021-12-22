package cluster

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

// GetClustersCommand represents command that lists all clusters for the account
type GetClustersCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetClustersCommand
func (command *GetClustersCommand) Execute() ([]rancherapiclient.Cluster, error) {

	server, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersGet(context.Background()).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("GetClustersCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return server, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("GetClustersCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewGetClustersCommand constructs new commmand of this type
func NewGetClustersCommand(receiver receiver.BMCSDK) *GetClustersCommand {

	return &GetClustersCommand{receiver}
}
