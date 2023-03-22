package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
)

// UpdateStorageNetworkVolumeCommand represents command that updates the storage network's volume details.
type UpdateStorageNetworkVolumeCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
	volumeId         string
	volumeUpdate     networkstorageapiclient.VolumeUpdate
}

// Execute runs UpdateStorageNetworkVolumeCommand
func (command *UpdateStorageNetworkVolumeCommand) Execute() (*networkstorageapiclient.Volume, error) {

	volume, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdPatch(context.Background(), command.storageNetworkId, command.volumeId).VolumeUpdate(command.volumeUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return volume, nil
	}
	return nil, fmt.Errorf("UpdateStorageNetworkVolumeCommand %s", errResolver.Error)
}

//NewUpdateStorageNetworkVolumeCommand constructs new commmand of this type
func NewUpdateStorageNetworkVolumeCommand(receiver receiver.BMCSDK, storageNetworkId string, volumeId string, volumeUpdate networkstorageapiclient.VolumeUpdate) *UpdateStorageNetworkVolumeCommand {

	return &UpdateStorageNetworkVolumeCommand{receiver, storageNetworkId, volumeId, volumeUpdate}
}
