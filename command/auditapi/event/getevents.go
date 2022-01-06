package event

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	auditapiclient "github.com/phoenixnap/go-sdk-bmc/auditapi"
)

// GetEventsCommand represents command that retrieves event logs
type GetEventsCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetEventsCommand
func (command *GetEventsCommand) Execute() ([]auditapiclient.Event, error) {

	events, httpResponse, err := command.receiver.AuditAPIClient.EventsApi.EventsGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return events, nil
	}
	return nil, fmt.Errorf("GetEventsCommand %s", errResolver.Error)
}

//NewGetEventsCommand constructs new commmand of this type
func NewGetEventsCommand(receiver receiver.BMCSDK) *GetEventsCommand {

	return &GetEventsCommand{receiver}
}
