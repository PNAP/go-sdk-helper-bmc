package event

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	auditapiclient "github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
)

// GetEventsCommand represents command that retrieves event logs
// Use NewGetEventsCommand to initialize command properly.
type GetEventsCommand struct {
	receiver receiver.BMCSDK
	query    dto.Query
}

// Execute runs GetEventsCommand
func (command *GetEventsCommand) Execute() ([]auditapiclient.Event, error) {
	from := command.query.From
	to := command.query.To
	limit := command.query.Limit
	order := command.query.Order
	username := command.query.Username
	verb := command.query.Verb
	uri := command.query.Uri

	x1 := command.receiver.AuditAPIClient.EventsAPI.EventsGet(context.Background())

	if !from.IsZero() {
		x1 = x1.From(from)
	}
	if !to.IsZero() {
		x1 = x1.To(to)
	}
	if limit != 0 {
		x1 = x1.Limit(limit)
	}
	if order != "" {
		x1 = x1.Order(order)
	}
	if username != "" {
		x1 = x1.Username(username)
	}
	if verb != "" {
		x1 = x1.Verb(verb)
	}
	if uri != "" {
		x1 = x1.Uri(uri)
	}

	events, httpResponse, err := x1.Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return events, nil
	}
	return nil, fmt.Errorf("GetEventsCommand %s", errResolver.Error)
}

//NewGetEventsCommand constructs new commmand of this type
func NewGetEventsCommand(receiver receiver.BMCSDK, query dto.Query) *GetEventsCommand {

	return &GetEventsCommand{receiver, query}
}
