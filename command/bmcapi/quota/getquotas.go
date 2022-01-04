package quota

import (
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"

	//"net/http"
	"context"

	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	bmcapiclient "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

// GetQuotasCommand represents command that lists all quotas for the account
type GetQuotasCommand struct {
	receiver receiver.BMCSDK
}

// Execute runs GetQuotasCommand
func (command *GetQuotasCommand) Execute() ([]bmcapiclient.Quota, error) {

	server, httpResponse, err := command.receiver.APIClient.QuotasApi.QuotasGet(context.Background()).Execute()

	if err != nil {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, err
		}
		return nil, fmt.Errorf("GetQuotasCommand Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)

		//return nil, err
	} else if httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return server, nil
	} else {
		response := &dto.ErrorMessage{}
		error := response.FromBytes(httpResponse)
		if error != nil {
			return nil, error
		}
		return nil, fmt.Errorf("API Returned Code %v Message: %s Validation Errors: %s", httpResponse.StatusCode, response.Message, response.ValidationErrors)
	}

}

//NewGetQuotasCommand constructs new command of this type
func NewGetQuotasCommand(receiver receiver.BMCSDK) *GetQuotasCommand {

	return &GetQuotasCommand{receiver}
}
