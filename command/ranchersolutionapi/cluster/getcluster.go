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

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return &cluster, nil
	}
	return nil, fmt.Errorf("GetClusterCommand %s", errResolver.Error)
}

//NewGetClusterCommand constructs new commmand of this type
func NewGetClusterCommand(receiver receiver.BMCSDK, clusterID string) *GetClusterCommand {

	return &GetClusterCommand{receiver, clusterID}
}
