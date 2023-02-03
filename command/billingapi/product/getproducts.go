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

	_, httpResponse, err := command.receiver.BillingAPIClient.ProductsApi.ProductsGet(context.Background()).ProductCode(productCode).ProductCategory(productCategory).SkuCode(skuCode).Location(location).Execute()

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
