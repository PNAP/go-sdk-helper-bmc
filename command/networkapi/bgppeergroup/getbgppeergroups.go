package bgppeergroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	networkapiclient "github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
)

// GetBgpPeerGroupsCommand represents command that retrieves all BGP Peer Groups owned by account
type GetBgpPeerGroupsCommand struct {
	receiver receiver.BMCSDK
	query    *dto.Query
}

// Execute runs GetBgpPeerGroupsCommand
func (command *GetBgpPeerGroupsCommand) Execute() ([]networkapiclient.BgpPeerGroup, error) {

	var bgpPeerGroups []networkapiclient.BgpPeerGroup
	var httpResponse *http.Response
	var err error

	if command.query != nil {

		location := command.query.LocationString

		bgpPeerGroups, httpResponse, err = command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsGet(context.Background()).Location(location).Execute()
	} else {

		bgpPeerGroups, httpResponse, err = command.receiver.NetworkAPIClient.BGPPeerGroupsAPI.BgpPeerGroupsGet(context.Background()).Execute()
	}

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return bgpPeerGroups, nil
	}
	return nil, fmt.Errorf("GetBgpPeerGroupsCommand %s", errResolver.Error)
}

//NewGetBgpPeerGroupsCommand constructs new commmand of this type
func NewGetBgpPeerGroupsCommand(receiver receiver.BMCSDK) *GetBgpPeerGroupsCommand {

	return &GetBgpPeerGroupsCommand{receiver, nil}
}

//NewGetBgpPeerGroupsWithQueryCommand constructs new commmand of this type
func NewGetBgpPeerGroupsWithQueryCommand(receiver receiver.BMCSDK, query *dto.Query) *GetBgpPeerGroupsCommand {

	return &GetBgpPeerGroupsCommand{receiver, query}
}
