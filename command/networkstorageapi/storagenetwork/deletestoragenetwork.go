package storagenetwork

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// DeleteStorageNetworkCommand represents command that deletes specific storage network and volume on the account
type DeleteStorageNetworkCommand struct {
	receiver         receiver.BMCSDK
	storagenetworkID string
}

// Execute runs DeleteStorageNetworkCommand
func (command *DeleteStorageNetworkCommand) Execute() error {

	httpResponse, err := command.receiver.NetworkStorageAPIClient.StorageNetworksApi.StorageNetworksIdDelete(context.Background(), command.storagenetworkID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return nil
	}
	return fmt.Errorf("DeleteStorageNetworkCommand %s", errResolver.Error)
}

//NewDeleteStorageNetworkCommand constructs new commmand of this type
func NewDeleteStorageNetworkCommand(receiver receiver.BMCSDK, storagenetworkID string) *DeleteStorageNetworkCommand {

	return &DeleteStorageNetworkCommand{receiver, storagenetworkID}
}
