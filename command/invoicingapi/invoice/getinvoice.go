package invoice

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	invoicingapiclient "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
)

// GetInvoiceCommand represents command that retrieves a specific invoice for the account.
// Use NewGetInvoiceCommand to initialize command properly.
type GetInvoiceCommand struct {
	receiver  receiver.BMCSDK
	invoiceID string
}

// Execute retrieves a specific invoice for the account.
func (command *GetInvoiceCommand) Execute() (*invoicingapiclient.Invoice, error) {

	invoice, httpResponse, err := command.receiver.InvoicingAPIClient.InvoicesAPI.InvoicesInvoiceIdGet(context.Background(), command.invoiceID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return invoice, nil
	}
	return nil, fmt.Errorf("GetInvoiceCommand %s", errResolver.Error)
}

//NewGetInvoiceCommand constructs new commmand of this type
func NewGetInvoiceCommand(receiver receiver.BMCSDK, invoiceID string) *GetInvoiceCommand {

	return &GetInvoiceCommand{receiver, invoiceID}
}
