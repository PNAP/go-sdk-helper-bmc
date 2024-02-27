package product

import (
	"context"
	"fmt"

	"github.com/PNAP/go-sdk-helper-bmc/dto"
	"github.com/PNAP/go-sdk-helper-bmc/receiver"
)

// GetProductsCommand represents command that retrieves products for the account.
// Use NewGetProductsCommand to initialize command properly.
type GetProductsCommand struct {
	receiver     receiver.BMCSDK
	productQuery dto.ProductQuery
}

// Execute retrieves products for the account.
func (command *GetProductsCommand) Execute() ([]dto.Product, error) {
	productCode := command.productQuery.ProductCode
	productCategory := command.productQuery.ProductCategory
	skuCode := command.productQuery.SKUCode
	location := command.productQuery.Location

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

	_, httpResponse, err := x1.Execute()

	errResolver := dto.NewErrorResolver(httpResponse, err)

	if errResolver.Error == nil {
		var productResponse = &dto.Products{}
		productResponse.FromBytes(httpResponse)
		respList := *productResponse
		return respList, nil
	}
	return nil, fmt.Errorf("GetProductsCommand %s", errResolver.Error)

}

//NewGetProductsCommand constructs new commmand of this type
func NewGetProductsCommand(reciever receiver.BMCSDK, productQuery dto.ProductQuery) *GetProductsCommand {

	return &GetProductsCommand{reciever, productQuery}
}
