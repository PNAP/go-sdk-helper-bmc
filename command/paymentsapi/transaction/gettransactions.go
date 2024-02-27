package transaction

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	paymentsapiclient "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
)

// GetTransactionsCommand represents command that retrieves a paginated list of transactions for the account.
// Use NewGetTransactionsCommand to initialize command properly.
type GetTransactionsCommand struct {
	receiver receiver.BMCSDK
	query    dto.Query
}

// Execute retrieves a paginated list of transactions for the account.
func (command *GetTransactionsCommand) Execute() (*paymentsapiclient.PaginatedTransactions, error) {
	limit := command.query.Limit
	offset := command.query.Offset
	sortDirection := command.query.SortDirection
	sortField := command.query.SortField
	from := command.query.From
	to := command.query.To

	x1 := command.receiver.PaymentsAPIClient.TransactionsAPI.TransactionsGet(context.Background())
	if limit != 0 {
		x1 = x1.Limit(limit)
	}
	if sortDirection != "" {
		x1 = x1.SortDirection(sortDirection)
	}
	if sortField != "" {
		x1 = x1.SortField(sortField)
	}
	if !from.IsZero() {
		x1 = x1.From(from)
	}
	if !to.IsZero() {
		x1 = x1.To(to)
	}

	transactions, httpResponse, err := x1.Offset(offset).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return transactions, nil
	}
	return nil, fmt.Errorf("GetTransactionsCommand %s", errResolver.Error)
}

//NewGetTransactionsCommand constructs new commmand of this type
func NewGetTransactionsCommand(receiver receiver.BMCSDK, query dto.Query) *GetTransactionsCommand {

	return &GetTransactionsCommand{receiver, query}
}
