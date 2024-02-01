package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
)

// RequestEditQuotaLimitCommand represents command that sends a request to edit the limit of a quota
type RequestEditQuotaLimitCommand struct {
	receiver              receiver.BMCSDK
	quotaID               string
	quotaEditLimitRequest bmcapiclient.QuotaEditLimitRequest
}

// Execute runs RequestEditQuotaLimitCommand
func (command *RequestEditQuotaLimitCommand) Execute() error {

	httpResponse, err := command.receiver.APIClient.QuotasAPI.QuotasQuotaIdActionsRequestEditPost(context.Background(), command.quotaID).QuotaEditLimitRequest(command.quotaEditLimitRequest).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return nil
	}
	return fmt.Errorf("RequestEditQuotaLimitCommand %s", errResolver.Error)
}

//NewRequestEditQuotaLimitCommand constructs new command of this type
func NewRequestEditQuotaLimitCommand(receiver receiver.BMCSDK, quotaID string, quotaEditLimitRequest bmcapiclient.QuotaEditLimitRequest) *RequestEditQuotaLimitCommand {

	return &RequestEditQuotaLimitCommand{receiver, quotaID, quotaEditLimitRequest}
}
