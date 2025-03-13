package bgppeergroup

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
)

// UpdateBgpPeerGroupCommand represents command that modifies a specific BGP Peer Group owned by account
type UpdateBgpPeerGroupCommand struct {
	receiver          receiver.BMCSDK
	bgpPeerGroupID    string
	bgpPeerGroupPatch networkapiclient.BgpPeerGroupPatch
}

// Execute runs UpdateBgpPeerGroupCommand
func (command *UpdateBgpPeerGroupCommand) Execute() (*networkapiclient.BgpPeerGroup, error) {

	bgpPeerGroup, httpResponse, err := command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdPatch(context.Background(), command.bgpPeerGroupID).
		BgpPeerGroupPatch(command.bgpPeerGroupPatch).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return bgpPeerGroup, nil
	}
	return nil, fmt.Errorf("UpdateBgpPeerGroupCommand %s", errResolver.Error)
}

//NewUpdateBgpPeerGroupCommand constructs new commmand of this type
func NewUpdateBgpPeerGroupCommand(receiver receiver.BMCSDK, bgpPeerGroupID string, bgpPeerGroupPatch networkapiclient.BgpPeerGroupPatch) *UpdateBgpPeerGroupCommand {

	return &UpdateBgpPeerGroupCommand{receiver, bgpPeerGroupID, bgpPeerGroupPatch}
}
