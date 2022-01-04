package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// RequestEditQuotaLimitCommand represents command that sends a request to edit the limit of a quota
type RequestEditQuotaLimitCommand struct {
	receiver              receiver.BMCSDK
	quotaID               string
	quotaEditLimitRequest bmcapiclient.QuotaEditLimitRequest
}

// Execute runs RequestEditQuotaLimitCommand
func (command *RequestEditQuotaLimitCommand) Execute() error {

	httpResponse, err := command.receiver.APIClient.QuotasApi.QuotasQuotaIdActionsRequestEditPost(context.Background(), command.quotaID).QuotaEditLimitRequest(command.quotaEditLimitRequest).Execute()

	if err != nil && httpResponse == nil {
		return err
	} else if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return err
		}
		return fmt.Errorf("RequestEditQuotaLimitCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return error
		}
		return fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewRequestEditQuotaLimitCommand constructs new command of this type
func NewRequestEditQuotaLimitCommand(receiver receiver.BMCSDK, quotaID string, quotaEditLimitRequest bmcapiclient.QuotaEditLimitRequest) *RequestEditQuotaLimitCommand {

	return &RequestEditQuotaLimitCommand{receiver, quotaID, quotaEditLimitRequest}
}
