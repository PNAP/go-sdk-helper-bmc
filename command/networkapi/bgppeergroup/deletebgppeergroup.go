package bgppeergroup

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// DeleteBgpPeerGroupCommand represents command that deletes a specific BGP Peer Group owned by account
type DeleteBgpPeerGroupCommand struct {
	receiver       receiver.BMCSDK
	bgpPeerGroupID string
}

// Execute runs DeleteBgpPeerGroupCommand
func (command *DeleteBgpPeerGroupCommand) Execute() (*networkapiclient.BgpPeerGroup, error) {

	bgpPeerGroup, httpResponse, err := command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdDelete(context.Background(), command.bgpPeerGroupID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return bgpPeerGroup, nil
	}
	return nil, fmt.Errorf("DeleteBgpPeerGroupCommand %s", errResolver.Error)
}

//NewDeleteBgpPeerGroupCommand constructs new commmand of this type
func NewDeleteBgpPeerGroupCommand(receiver receiver.BMCSDK, bgpPeerGroupID string) *DeleteBgpPeerGroupCommand {

	return &DeleteBgpPeerGroupCommand{receiver, bgpPeerGroupID}
}
