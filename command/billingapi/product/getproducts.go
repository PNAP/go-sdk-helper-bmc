package product

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
	billingapiclient "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
)

// GetProductsCommand represents command that retrieves products for the account.
// Use NewGetProductsCommand to initialize command properly.
type GetProductsCommand struct {
	receiver receiver.BMCSDK
	query    *dto.Query
}

// Execute retrieves products for the account.
func (command *GetProductsCommand) Execute() ([]billingapiclient.ProductsGet200ResponseInner, error) {
	productCode := command.query.ProductCode
	productCategory := command.query.ProductCategory
	skuCode := command.query.SKUCode
	location := command.query.Location

	x1 := command.receiver.BillingAPIClient.ProductsAPI.ProductsGet(context.Background())

	if productCode != "" {
		x1 = x1.ProductCode(productCode)
	}
	if productCategory != "" {
		x1 = x1.ProductCategory(productCategory)
	}
	if skuCode != "" {
		x1 = x1.SkuCode(skuCode)
	}
	if location != "" {
		x1 = x1.Location(location)
	}

	products, httpResponse, err := x1.Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		return products, nil
	}
	return nil, fmt.Errorf("GetProductsCommand %s", errResolver.Error)

}

//NewGetProductsCommand constructs new commmand of this type
func NewGetProductsCommand(reciever receiver.BMCSDK, query *dto.Query) *GetProductsCommand {

	return &GetProductsCommand{reciever, query}
}
