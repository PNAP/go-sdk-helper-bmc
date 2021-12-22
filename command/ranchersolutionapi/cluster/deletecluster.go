package cluster

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

// DeleteClusterCommand represents command that deletes a cluster
type DeleteClusterCommand struct {
	receiver  receiver.BMCSDK
	clusterID string
}

// Execute runs DeleteClusterCommand
func (command *DeleteClusterCommand) Execute() (*rancherapiclient.DeleteResult, error) {

	result, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersIdDelete(context.Background(), command.clusterID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("DeleteClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return &result, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("DeleteClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewDeleteClusterCommand constructs new commmand of this type
func NewDeleteClusterCommand(receiver receiver.BMCSDK, clusterID string) *DeleteClusterCommand {

	return &DeleteClusterCommand{receiver, clusterID}
}
