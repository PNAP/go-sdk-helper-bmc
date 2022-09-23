package cluster

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	rancherapiclient "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
)

// GetClustersCommand represents command that lists all clusters for the account
type GetClustersCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetClustersCommand
func (command *GetClustersCommand) Execute() ([]rancherapiclient.Cluster, error) {

	clusters, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return clusters, nil
	}
	return nil, fmt.Errorf("GetClustersCommand %s", errResolver.Error)
}

//NewGetClustersCommand constructs new commmand of this type
func NewGetClustersCommand(receiver receiver.BMCSDK) *GetClustersCommand {

	return &GetClustersCommand{receiver}
}
