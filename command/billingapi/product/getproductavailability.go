package product

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// GetProductAvailabilityCommand represents command that retrieves product availabilities for the account.
// Use NewGetProductAvailabilityCommand to initialize command properly.
type GetProductAvailabilityCommand struct {
	receiver                 receiver.BMCSDK
	productAvailabilityQuery dto.ProductAvailabilityQuery
}

// Execute retrieves product availabilities for the account.
func (command *GetProductAvailabilityCommand) Execute() ([]dto.ProductAvailability, error) {

	productCategory := command.productAvailabilityQuery.ProductCategory
	productCode := command.productAvailabilityQuery.ProductCode
	showOnlyMinQuantityAvailable := command.productAvailabilityQuery.ShowOnlyMinQuantityAvailable

	loc := command.productAvailabilityQuery.Location
	location := make([]billingapi.LocationEnum, len(loc))
	for i, j := range loc {
		location[i] = billingapi.LocationEnum(j)
	}

	solution := command.productAvailabilityQuery.Solution
	minQuantity := command.productAvailabilityQuery.MinQuantity

	x1 := command.receiver.BillingAPIClient.ProductAvailabilityAPI.ProductAvailabilityGet(context.Background())

	if productCategory != nil {
		x1 = x1.ProductCategory(productCategory)
	}
	if productCode != nil {
		x1 = x1.ProductCode(productCode)
	}
	if location != nil {
		x1 = x1.Location(location)
	}
	if solution != nil {
		x1 = x1.Solution(solution)
	}
	if minQuantity != 0 {
		x1 = x1.MinQuantity(minQuantity)
	}

	_, httpResponse, err := x1.ShowOnlyMinQuantityAvailable(showOnlyMinQuantityAvailable).Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var productResponse = &dto.ProductAvailabilities{}
		productResponse.FromBytes(httpResponse)
		respList := *productResponse
		return respList, nil
	}
	return nil, fmt.Errorf("GetProductAvailabilityCommand %s", errResolver.Error)

}

//NewGetProductAvailabilityCommand constructs new commmand of this type
func NewGetProductAvailabilityCommand(requester receiver.BMCSDK, productAvailabilityQuery dto.ProductAvailabilityQuery) *GetProductAvailabilityCommand {

	return &GetProductAvailabilityCommand{requester, productAvailabilityQuery}
}
