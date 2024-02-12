package invoice

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	invoicingapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

// GetInvoicesCommand represents command that retrieves invoices for the account.
// Use NewGetInvoicesCommand to initialize command properly.
type GetInvoicesCommand struct {
	receiver receiver.BMCSDK
	query    dto.Query
}

// Execute retrieves invoices for the account.
func (command *GetInvoicesCommand) Execute() (*invoicingapiclient.PaginatedInvoices, error) {
	number := command.query.Number
	status := command.query.Status
	sentOnFrom := command.query.SentOnFrom
	sentOnTo := command.query.SentOnTo
	limit := command.query.Limit
	offset := command.query.Offset
	sortField := command.query.SortField
	sortDirection := command.query.SortDirection

	x1 := command.receiver.InvoicingAPIClient.InvoicesAPI.InvoicesGet(context.Background())
	if number != "" {
		x1 = x1.Number(number)
	}
	if status != "" {
		x1 = x1.Status(status)
	}
	if !sentOnFrom.IsZero() {
		x1 = x1.SentOnFrom(sentOnFrom)
	}
	if !sentOnTo.IsZero() {
		x1 = x1.SentOnTo(sentOnTo)
	}
	if limit != 0 {
		x1 = x1.Limit(limit)
	}
	if sortField != "" {
		x1 = x1.SortField(sortField)
	}
	if sortDirection != "" {
		x1 = x1.SortDirection(sortDirection)
	}

	invoices, httpResponse, err := x1.Offset(offset).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return invoices, nil
	}
	return nil, fmt.Errorf("GetInvoicesCommand %s", errResolver.Error)
}

//NewGetInvoicesCommand constructs new commmand of this type
func NewGetInvoicesCommand(receiver receiver.BMCSDK, query dto.Query) *GetInvoicesCommand {

	return &GetInvoicesCommand{receiver, query}
}
