package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// CreateStorageNetworkVolumeCommand represents command that creates a volume on specific storage network on the account
type CreateStorageNetworkVolumeCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
	volumeCreate     networkstorageapiclient.VolumeCreate
}

// Execute runs CreateStorageNetworkVolumeCommand
func (command *CreateStorageNetworkVolumeCommand) Execute() (*networkstorageapiclient.Volume, error) {

	volume, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesPost(context.Background(), command.storageNetworkId).
		VolumeCreate(command.volumeCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return volume, nil
	}
	return nil, fmt.Errorf("CreateStorageNetworkVolumeCommand %s", errResolver.Error)
}

//NewCreateStorageNetworkVolumeCommand constructs new commmand of this type
func NewCreateStorageNetworkVolumeCommand(receiver receiver.BMCSDK, storageNetworkId string, volumeCreate networkstorageapiclient.VolumeCreate) *CreateStorageNetworkVolumeCommand {

	return &CreateStorageNetworkVolumeCommand{receiver, storageNetworkId, volumeCreate}
}
