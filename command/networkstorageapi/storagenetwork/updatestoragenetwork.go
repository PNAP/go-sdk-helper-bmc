package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// UpdateStorageNetworkCommand represents command that updates storage network details
type UpdateStorageNetworkCommand struct {
	receiver             receiver.BMCSDK
	storageNetworkID     string
	storageNetworkUpdate networkstorageapiclient.StorageNetworkUpdate
}

// Execute runs UpdateStorageNetworkCommand
func (command *UpdateStorageNetworkCommand) Execute() (*networkstorageapiclient.StorageNetwork, error) {

	storageNetwork, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksIdPatch(context.Background(), command.storageNetworkID).StorageNetworkUpdate(command.storageNetworkUpdate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return storageNetwork, nil
	}
	return nil, fmt.Errorf("UpdateStorageNetworkCommand %s", errResolver.Error)
}

//NewUpdateStorageNetworkCommand constructs new commmand of this type
func NewUpdateStorageNetworkCommand(receiver receiver.BMCSDK, storageNetworkID string, storageNetworkUpdate networkstorageapiclient.StorageNetworkUpdate) *UpdateStorageNetworkCommand {

	return &UpdateStorageNetworkCommand{receiver, storageNetworkID, storageNetworkUpdate}
}
