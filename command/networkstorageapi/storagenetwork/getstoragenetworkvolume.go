package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// GetStorageNetworkVolumeCommand represents command that retrieves details about specific volume on a storage network on the account
type GetStorageNetworkVolumeCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
	volumeId         string
}

// Execute runs GetStorageNetworkVolumeCommand
func (command *GetStorageNetworkVolumeCommand) Execute() (*networkstorageapiclient.Volume, error) {

	volume, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdGet(context.Background(), command.storageNetworkId, command.volumeId).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return volume, nil
	}
	return nil, fmt.Errorf("GetStorageNetworkVolumeCommand %s", errResolver.Error)
}

//NewGetStorageNetworkVolumeCommand constructs new commmand of this type
func NewGetStorageNetworkVolumeCommand(receiver receiver.BMCSDK, storageNetworkId string, volumeId string) *GetStorageNetworkVolumeCommand {

	return &GetStorageNetworkVolumeCommand{receiver, storageNetworkId, volumeId}
}
