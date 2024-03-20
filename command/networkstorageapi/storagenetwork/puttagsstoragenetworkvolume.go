package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// PutTagsStorageNetworkVolumeCommand represents command that overwrites tags assigned for specific volume on a storage network on the account
type PutTagsStorageNetworkVolumeCommand struct {
	receiver             receiver.BMCSDK
	storageNetworkId     string
	volumeId             string
	tagAssignmentRequest []networkstorageapiclient.TagAssignmentRequest
}

// Execute runs PutTagsStorageNetworkVolumeCommand
func (command *PutTagsStorageNetworkVolumeCommand) Execute() (*networkstorageapiclient.Volume, error) {

	volume, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksStorageNetworkIdVolumesVolumeIdTagsPut(context.Background(), command.storageNetworkId, command.volumeId).TagAssignmentRequest(command.tagAssignmentRequest).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return volume, nil
	}
	return nil, fmt.Errorf("PutTagsStorageNetworkVolumeCommand %s", errResolver.Error)
}

//NewPutTagsStorageNetworkVolumeCommand constructs new commmand of this type
func NewPutTagsStorageNetworkVolumeCommand(receiver receiver.BMCSDK, storageNetworkId string, volumeId string, tagAssignmentRequest []networkstorageapiclient.TagAssignmentRequest) *PutTagsStorageNetworkVolumeCommand {

	return &PutTagsStorageNetworkVolumeCommand{receiver, storageNetworkId, volumeId, tagAssignmentRequest}
}
