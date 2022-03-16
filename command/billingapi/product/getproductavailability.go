package product

import (
	"fmt"
	"strconv"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// GetProductAvailabilityCommand represents command that retrieves product availabilities for the account.
// Use NewGetProductAvailabilityCommand to initialize command properly.
type GetProductAvailabilityCommand struct {
	receiver                 receiver.BMCSDK
	productAvailabilityQuery dto.ProductAvailabilityQuery
}

// Execute retrieves product availabilities for the account.
func (command *GetProductAvailabilityCommand) Execute() ([]dto.ProductAvailability, error) {
	var req = command.receiver
	var apiPrefix = "billing/v1/product-availability?"

	var queryParams = ""
	productCategory := command.productAvailabilityQuery.ProductCategory
	if productCategory != nil {
		proCat := *productCategory
		if len(proCat) > 0 {
			for _, j := range proCat {
				if j != "" {
					queryParams = queryParams + "&productCategory=" + j
				}
			}
		}
	}
	productCode := command.productAvailabilityQuery.ProductCode
	if productCode != nil {
		proCode := *productCode
		if len(proCode) > 0 {
			for _, j := range proCode {
				if j != "" {
					queryParams = queryParams + "&productCode=" + j
				}
			}
		}
	}
	showOnlyMinQuantityAvailable := command.productAvailabilityQuery.ShowOnlyMinQuantityAvailable
	if showOnlyMinQuantityAvailable != nil {
		somqaa := *showOnlyMinQuantityAvailable
		somqab := strconv.FormatBool(somqaa)
		queryParams = queryParams + "&showOnlyMinQuantityAvailable=" + somqab
	}
	location := command.productAvailabilityQuery.Location
	if location != nil {
		loc := *location
		if len(loc) > 0 {
			for _, j := range loc {
				if j != "" {
					queryParams = queryParams + "&location=" + j
				}
			}
		}
	}
	solution := command.productAvailabilityQuery.Solution
	if solution != nil {
		sol := *solution
		if len(sol) > 0 {
			for _, j := range sol {
				if j != "" {
					queryParams = queryParams + "&solution=" + j
				}
			}
		}
	}
	minQuantity := command.productAvailabilityQuery.MinQuantity
	if minQuantity != nil {
		mqa := *minQuantity
		mqs := strconv.FormatInt(int64(mqa), 10)
		if mqs != "" {
			queryParams = queryParams + "&minQuantity=" + mqs
		}
	}

	httpResponse, err := req.PNAPClient.Get(apiPrefix + queryParams)

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
