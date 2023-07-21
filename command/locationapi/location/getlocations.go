package location

import (
	"fmt"

	"context"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	locationapiclient "github.com/phoenixnap/go-sdk-bmc/locationapi"
)

// GetLocationsCommand represents command that retrieves the locations info
type GetLocationsCommand struct {
	receiver receiver.BMCSDK
	query    dto.Query
}

// Execute runs GetLocationsCommand
func (command *GetLocationsCommand) Execute() ([]locationapiclient.Location, error) {

	location := command.query.Location
	productCategory := command.query.ProductCategory

	locations, httpResponse, err := command.receiver.LocationAPIClient.LocationsApi.GetLocations(context.Background()).Location(location).ProductCategory(productCategory).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return locations, nil
	}
	return nil, fmt.Errorf("GetLocationsCommand %s", errResolver.Error)
}

//NewGetLocationsCommand constructs new commmand of this type
func NewGetLocationsCommand(receiver receiver.BMCSDK, query dto.Query) *GetLocationsCommand {

	return &GetLocationsCommand{receiver, query}
}
