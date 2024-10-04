package dto

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
)

//PricingPlan represents pricing plan details
type PricingPlan struct {
	SKU                   string               `json:"sku"`
	SKUDescription        string               `json:"skuDescription,omitempty"`
	Location              string               `json:"location"`
	PricingModel          string               `json:"pricingModel"`
	Price                 float32              `json:"price"`
	PriceUnit             string               `json:"priceUnit"`
	ApplicableDiscounts   *ApplicableDiscounts `json:"applicableDiscounts,omitempty"`
	CorrelatedProductCode string               `json:"correlatedProductCode,omitempty"`
	PackageQuantity       int32                `json:"packageQuantity,omitempty"`
	PackageUnit           string               `json:"packageUnit,omitempty"`
}

//FromBytes performs conversion of http response to the representing struct
func (dto *PricingPlan) FromBytes(resp *http.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, dto)
	}
	return err
}
