package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetQuotaCommand represents command that retrieves a quota
type GetQuotaCommand struct {
	receiver receiver.BMCSDK
	quotaID  string
}

// Execute runs GetQuotaCommand
func (command *GetQuotaCommand) Execute() (*bmcapiclient.Quota, error) {

	quota, httpResponse, err := command.receiver.APIClient.QuotasApi.QuotasQuotaIdGet(context.Background(), command.quotaID).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("GetQuotaCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return &quota, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewGetQuotaCommand constructs new command of this type
func NewGetQuotaCommand(receiver receiver.BMCSDK, quotaID string) *GetQuotaCommand {

	return &GetQuotaCommand{receiver, quotaID}
}
