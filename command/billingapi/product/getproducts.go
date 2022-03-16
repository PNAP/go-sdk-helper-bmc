package product

import (
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
	var req = command.receiver
	var apiPrefix = "billing/v1/products?"

	var queryParams = ""
	productCode := command.productQuery.ProductCode
	if productCode != "" {
		queryParams = queryParams + "&productCode=" + productCode
	}
	productCategory := command.productQuery.ProductCategory
	if productCategory != "" {
		queryParams = queryParams + "&productCategory=" + productCategory
	}
	skuCode := command.productQuery.SKUCode
	if skuCode != "" {
		queryParams = queryParams + "&skuCode=" + skuCode
	}
	location := command.productQuery.Location
	if location != "" {
		queryParams = queryParams + "&location=" + location
	}

	httpResponse, err := req.PNAPClient.Get(apiPrefix + queryParams)

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
func NewGetProductsCommand(requester receiver.BMCSDK, productQuery dto.ProductQuery) *GetProductsCommand {

	return &GetProductsCommand{requester, productQuery}
}
