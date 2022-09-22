package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
)

// CreateStorageNetworkCommand represents command that creates a storage network and volume on the account
type CreateStorageNetworkCommand struct {
	receiver             receiver.BMCSDK
	storageNetworkCreate networkstorageapiclient.StorageNetworkCreate
}

// Execute runs CreateStorageNetworksCommand
func (command *CreateStorageNetworkCommand) Execute() (*networkstorageapiclient.StorageNetwork, error) {

	storageNetwork, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksApi.StorageNetworksPost(context.Background()).StorageNetworkCreate(command.storageNetworkCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return storageNetwork, nil
	}
	return nil, fmt.Errorf("CreateStorageNetworkCommand %s", errResolver.Error)
}

//NewCreateStorageNetworksCommand constructs new commmand of this type
func NewCreateStorageNetworkCommand(receiver receiver.BMCSDK, storageNetworkCreate networkstorageapiclient.StorageNetworkCreate) *CreateStorageNetworkCommand {

	return &CreateStorageNetworkCommand{receiver, storageNetworkCreate}
}
