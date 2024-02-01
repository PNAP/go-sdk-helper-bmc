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

	invoices, httpResponse, err := command.receiver.InvoicingAPIClient.InvoicesAPI.InvoicesGet(context.Background()).Number(number).Status(status).SentOnFrom(sentOnFrom).
		SentOnTo(sentOnTo).Limit(limit).Offset(offset).SortField(sortField).SortDirection(sortDirection).Execute()

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
