package cluster

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

// CreateClusterCommand represents command that adds new Rancher Server Deployment to the account
type CreateClusterCommand struct {
	receiver receiver.BMCSDK
	cluster  rancherapiclient.Cluster
}

// Execute runs CreateClusterCommand
func (command *CreateClusterCommand) Execute() (*rancherapiclient.Cluster, error) {

	cluster, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersPost(context.Background()).Cluster(command.cluster).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return cluster, nil
	}
	return nil, fmt.Errorf("CreateClusterCommand %s", errResolver.Error)
}

//NewCreateClusterCommand constructs new commmand of this type
func NewCreateClusterCommand(receiver receiver.BMCSDK, cluster rancherapiclient.Cluster) *CreateClusterCommand {

	return &CreateClusterCommand{receiver, cluster}
}
