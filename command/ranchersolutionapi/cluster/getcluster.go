package cluster

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

// GetClusterCommand represents command that retrieves a cluster
type GetClusterCommand struct {
	receiver  receiver.BMCSDK
	clusterID string
}

// Execute runs GetClusterCommand
func (command *GetClusterCommand) Execute() (*rancherapiclient.Cluster, error) {

	cluster, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersIdGet(context.Background(), command.clusterID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("GetClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return &cluster, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("GetClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewGetClusterCommand constructs new commmand of this type
func NewGetClusterCommand(receiver receiver.BMCSDK, clusterID string) *GetClusterCommand {

	return &GetClusterCommand{receiver, clusterID}
}
