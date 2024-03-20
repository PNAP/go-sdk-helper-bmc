package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// GetStorageNetworkVolumesCommand represents command that lists all volumes belonging to specific storage network
type GetStorageNetworkVolumesCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
}

// Execute runs GetStorageNetworkVolumesCommand
func (command *GetStorageNetworkVolumesCommand) Execute() ([]networkstorageapiclient.Volume, error) {

	volumes, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesGet(context.Background(), command.storageNetworkId).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return volumes, nil
	}
	return nil, fmt.Errorf("GetStorageNetworkVolumesCommand %s", errResolver.Error)
}

//NewGetStorageNetworkVolumesCommand constructs new commmand of this type
func NewGetStorageNetworkVolumesCommand(receiver receiver.BMCSDK, storageNetworkId string) *GetStorageNetworkVolumesCommand {

	return &GetStorageNetworkVolumesCommand{receiver, storageNetworkId}
}
