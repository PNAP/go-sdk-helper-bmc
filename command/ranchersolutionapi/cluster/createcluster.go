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

	server, httpResponse, err := command.receiver.RancherAPIClient.ClustersApi.ClustersPost(context.Background()).Cluster(command.cluster).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("CreateClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return &server, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("CreateClusterCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewCreateClusterCommand constructs new commmand of this type
func NewCreateClusterCommand(receiver receiver.BMCSDK, cluster rancherapiclient.Cluster) *CreateClusterCommand {

	return &CreateClusterCommand{receiver, cluster}
}