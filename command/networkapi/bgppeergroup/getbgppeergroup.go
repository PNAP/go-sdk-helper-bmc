package bgppeergroup

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// GetBgpPeerGroupCommand represents command that retrieves a specific BGP Peer Group owned by account
type GetBgpPeerGroupCommand struct {
	receiver       receiver.BMCSDK
	bgpPeerGroupID string
}

// Execute runs GetBgpPeerGroupCommand
func (command *GetBgpPeerGroupCommand) Execute() (*networkapiclient.BgpPeerGroup, error) {

	bgpPeerGroup, httpResponse, err := command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsPeerGroupIdGet(context.Background(), command.bgpPeerGroupID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return bgpPeerGroup, nil
	}
	return nil, fmt.Errorf("GetBgpPeerGroupCommand %s", errResolver.Error)
}

//NewGetBgpPeerGroupCommand constructs new commmand of this type
func NewGetBgpPeerGroupCommand(receiver receiver.BMCSDK, bgpPeerGroupID string) *GetBgpPeerGroupCommand {

	return &GetBgpPeerGroupCommand{receiver, bgpPeerGroupID}
}
