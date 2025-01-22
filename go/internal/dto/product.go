package dto

import (
	"github.com/JordanMarcelino/widatech-technical/internal/entity"
	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID               int64           `json:"id"`
	InvoiceNo        string          `json:"invoice_no"`
	Name             string          `json:"name"`
	Quantity         int64           `json:"quantity"`
	TotalCostOfGoods decimal.Decimal `json:"total_cost_of_goods"`
	TotalPriceSold   decimal.Decimal `json:"total_price_sold"`
}

// TODO: add validation > 0
type CreateProductRequest struct {
	Name             string          `json:"name" binding:"required,min=2,max=255"`
	Quantity         int64           `json:"quantity" binding:"required,min=1"`
	TotalCostOfGoods decimal.Decimal `json:"total_cost_of_goods" binding:"required"`
	TotalPriceSold   decimal.Decimal `json:"total_price_sold" binding:"required"`
}

type UpdateProductRequest struct {
	Name             string          `json:"name" binding:"required,min=2,max=255"`
	Quantity         int64           `json:"quantity" binding:"required,min=1"`
	TotalCostOfGoods decimal.Decimal `json:"total_cost_of_goods" binding:"required,dgt=0"`
	TotalPriceSold   decimal.Decimal `json:"total_price_sold" binding:"required,dgt=0"`
}

func ToProductResponse(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		ID:               product.ID,
		InvoiceNo:        product.InvoiceNo,
		Name:             product.Name,
		Quantity:         product.Quantity,
		TotalCostOfGoods: product.TotalCostOfGoods,
		TotalPriceSold:   product.TotalPriceSold,
	}
}

func ToProductResponses(products []*entity.Product) []*ProductResponse {
	res := []*ProductResponse{}

	for _, product := range products {
		res = append(res, ToProductResponse(product))
	}

	return res
}
