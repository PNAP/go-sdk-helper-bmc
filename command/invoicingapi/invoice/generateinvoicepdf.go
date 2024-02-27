package invoice

import (
	"context"
	"fmt"
	"os"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// GenerateInvoicePdfCommand represents command that generates invoice details as PDF for a specific invoice on the account.
// Use NewGenerateInvoicePdfCommand to initialize command properly.
type GenerateInvoicePdfCommand struct {
	receiver  receiver.BMCSDK
	invoiceID string
}

// Execute generates invoice details as PDF for a specific invoice on the account.
func (command *GenerateInvoicePdfCommand) Execute() (*os.File, error) {

	file, httpResponse, err := command.receiver.InvoicingAPIClient.InvoicesAPI.InvoicesInvoiceIdGeneratePdfPost(context.Background(), command.invoiceID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return file, nil
	}
	return nil, fmt.Errorf("GenerateInvoicePdfCommand %s", errResolver.Error)
}

//NewGenerateInvoicePdfCommand constructs new commmand of this type
func NewGenerateInvoicePdfCommand(receiver receiver.BMCSDK, invoiceID string) *GenerateInvoicePdfCommand {

	return &GenerateInvoicePdfCommand{receiver, invoiceID}
}
