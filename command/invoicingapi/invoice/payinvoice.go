package invoice

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// PayInvoiceCommand represents command that manually pays a specific invoice for the account.
// Use NewPayInvoiceCommand to initialize command properly.
type PayInvoiceCommand struct {
	receiver  receiver.BMCSDK
	invoiceID string
	body      map[string]interface{}
}

// Execute manually pays a specific invoice for the account.
func (command *PayInvoiceCommand) Execute() (map[string]interface{}, error) {

	resp, httpResponse, err := command.receiver.InvoicingAPIClient.InvoicesAPI.InvoicesInvoiceIdPayPost(context.Background(), command.invoiceID).Body(command.body).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return resp, nil
	}
	return nil, fmt.Errorf("PayInvoiceCommand %s", errResolver.Error)
}

//NewPayInvoiceCommand constructs new commmand of this type
func NewPayInvoiceCommand(receiver receiver.BMCSDK, invoiceID string, body map[string]interface{}) *PayInvoiceCommand {

	return &PayInvoiceCommand{receiver, invoiceID, body}
}
