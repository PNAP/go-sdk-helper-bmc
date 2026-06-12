package server

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// UpdateServerIPXECommand represents command that updates the iPXE OS configuration for specific server.
type UpdateServerIPXECommand struct {
	receiver            receiver.BMCSDK
	serverID            string
	osConfigurationIPXE bmcapiclient.OsConfigurationIPXE
}

// Execute updates the iPXE OS configuration for specific server
func (command *UpdateServerIPXECommand) Execute() (*bmcapiclient.OsConfigurationIPXE, error) {

	ipxe, httpResponse, err := command.receiver.APIClient.ServersAPI.ServersServerIdOsConfigurationIpxePut(context.Background(), command.serverID).
		OsConfigurationIPXE(command.osConfigurationIPXE).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return ipxe, nil
	}
	return nil, fmt.Errorf("UpdateServerIPXECommand %s", errResolver.Error)
}

// NewUpdateServerIPXECommand constructs new commmand of this type
func NewUpdateServerIPXECommand(receiver receiver.BMCSDK, serverID string, osConfigurationIPXE bmcapiclient.OsConfigurationIPXE) *UpdateServerIPXECommand {

	return &UpdateServerIPXECommand{receiver, serverID, osConfigurationIPXE}
}
