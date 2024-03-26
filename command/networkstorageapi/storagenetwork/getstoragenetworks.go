package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkstorageapiclient "github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
)

// GetStorageNetworksCommand represents command that lists all storage networks on the account
type GetStorageNetworksCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetStorageNetworksCommand
func (command *GetStorageNetworksCommand) Execute() ([]networkstorageapiclient.StorageNetwork, error) {

	storageNetworks, httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksAPI.StorageNetworksGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return storageNetworks, nil
	}
	return nil, fmt.Errorf("GetStorageNetworksCommand %s", errResolver.Error)
}

//NewGetStorageNetworksCommand constructs new commmand of this type
func NewGetStorageNetworksCommand(receiver receiver.BMCSDK) *GetStorageNetworksCommand {

	return &GetStorageNetworksCommand{receiver}
}
