package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

// GetQuotasCommand represents command that lists all quotas for the account
type GetQuotasCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetQuotasCommand
func (command *GetQuotasCommand) Execute() ([]bmcapiclient.Quota, error) {

	quota, httpResponse, err := command.receiver.APIClient.QuotasApi.QuotasGet(context.Background()).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return quota, nil
	}
	return nil, fmt.Errorf("GetQuotasCommand %s", errResolver.Error)
}

//NewGetQuotasCommand constructs new command of this type
func NewGetQuotasCommand(receiver receiver.BMCSDK) *GetQuotasCommand {

	return &GetQuotasCommand{receiver}
}
