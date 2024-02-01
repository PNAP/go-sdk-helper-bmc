package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// GetQuotaCommand represents command that retrieves a quota
type GetQuotaCommand struct {
	receiver receiver.BMCSDK
	quotaID  string
}

// Execute runs GetQuotaCommand
func (command *GetQuotaCommand) Execute() (*bmcapiclient.Quota, error) {

	quota, httpResponse, err := command.receiver.APIClient.QuotasAPI.QuotasQuotaIdGet(context.Background(), command.quotaID).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return quota, nil
	}
	return nil, fmt.Errorf("GetQuotaCommand %s", errResolver.Error)
}

//NewGetQuotaCommand constructs new command of this type
func NewGetQuotaCommand(receiver receiver.BMCSDK, quotaID string) *GetQuotaCommand {

	return &GetQuotaCommand{receiver, quotaID}
}
