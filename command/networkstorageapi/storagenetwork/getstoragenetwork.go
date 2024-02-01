package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v2"
)

// GetStorageNetworkCommand represents command that retrieves details about specific storage network on the account
type GetStorageNetworkCommand struct {
	receiver         receiver.BMCSDK
	storageNetworkId string
}

// Execute runs GetStorageNetworkCommand
func (command *GetStorageNetworkCommand) Execute() (*networkstorageapiclient.StorageNetwork, error) {

	storageNetwork, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksIdGet(context.Background(), command.storageNetworkId).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return storageNetwork, nil
	}
	return nil, fmt.Errorf("GetStorageNetworkCommand %s", errResolver.Error)
}

//NewGetStorageNetworkCommand constructs new commmand of this type
func NewGetStorageNetworkCommand(receiver receiver.BMCSDK, storageNetworkId string) *GetStorageNetworkCommand {

	return &GetStorageNetworkCommand{receiver, storageNetworkId}
}
