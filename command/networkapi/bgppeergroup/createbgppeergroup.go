package bgppeergroup

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
)

// CreateBgpPeerGroupCommand represents command that configures a new BGP Peer Group
type CreateBgpPeerGroupCommand struct {
	receiver           receiver.BMCSDK
	bgpPeerGroupCreate networkapiclient.BgpPeerGroupCreate
}

// Execute runs CreateBgpPeerGroupCommand
func (command *CreateBgpPeerGroupCommand) Execute() (*networkapiclient.BgpPeerGroup, error) {

	bgpPeerGroup, httpResponse, err := command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsPost(context.Background()).
		BgpPeerGroupCreate(command.bgpPeerGroupCreate).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return bgpPeerGroup, nil
	}
	return nil, fmt.Errorf("CreateBgpPeerGroupCommand %s", errResolver.Error)
}

//NewCreateBgpPeerGroupCommand constructs new commmand of this type
func NewCreateBgpPeerGroupCommand(receiver receiver.BMCSDK, bgpPeerGroupCreate networkapiclient.BgpPeerGroupCreate) *CreateBgpPeerGroupCommand {

	return &CreateBgpPeerGroupCommand{receiver, bgpPeerGroupCreate}
}
