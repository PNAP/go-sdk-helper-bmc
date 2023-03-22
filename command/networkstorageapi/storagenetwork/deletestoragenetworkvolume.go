package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeleteStorageNetworkVolumeCommand represents command that deletes a volume on specific storage network on the account
type DeleteStorageNetworkVolumeCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
	volumeId         string
}

// Execute runs DeleteStorageNetworkVolumeCommand
func (command *DeleteStorageNetworkVolumeCommand) Execute() error {

	httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksApi.StorageNetworksStorageNetworkIdVolumesVolumeIdDelete(context.Background(), command.storageNetworkId, command.volumeId).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return nil
	}
	return fmt.Errorf("DeleteStorageNetworkVolumeCommand %s", errResolver.Error)
}

//NewDeleteStorageNetworkVolumeCommand constructs new commmand of this type
func NewDeleteStorageNetworkVolumeCommand(receiver receiver.BMCSDK, storageNetworkId string, volumeId string) *DeleteStorageNetworkVolumeCommand {

	return &DeleteStorageNetworkVolumeCommand{receiver, storageNetworkId, volumeId}
}
