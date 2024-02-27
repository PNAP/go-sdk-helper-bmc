package transaction

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	paymentsapiclient "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
)

// GetTransactionCommand represents command that retrieves a specific transaction for the account.
// Use NewGetTransactionCommand to initialize command properly.
type GetTransactionCommand struct {
	receiver      receiver.BMCSDK
	transactionID string
}

// Execute retrieves a specific transaction for the account.
func (command *GetTransactionCommand) Execute() (*paymentsapiclient.Transaction, error) {

	transaction, httpResponse, err := command.receiver.PaymentsAPIClient.TransactionsAPI.TransactionIdGet(context.Background(), command.transactionID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return transaction, nil
	}
	return nil, fmt.Errorf("GetTransactionCommand %s", errResolver.Error)
}

//NewGetTransactionCommand constructs new commmand of this type
func NewGetTransactionCommand(receiver receiver.BMCSDK, transactionID string) *GetTransactionCommand {

	return &GetTransactionCommand{receiver, transactionID}
}
