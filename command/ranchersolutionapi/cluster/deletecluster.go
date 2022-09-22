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

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return result, nil
	}
	return nil, fmt.Errorf("DeleteClusterCommand %s", errResolver.Error)
}

//NewDeleteClusterCommand constructs new commmand of this type
func NewDeleteClusterCommand(receiver receiver.BMCSDK, clusterID string) *DeleteClusterCommand {

	return &DeleteClusterCommand{receiver, clusterID}
}
