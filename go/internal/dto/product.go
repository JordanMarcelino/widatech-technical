package dto

import "github.com/shopspring/decimal"

type ProductResponse struct {
	ID               int64
	InvoiceNo        string
	Name             string
	Quantity         int64
	TotalCostOfGoods decimal.Decimal
	TotalPriceSold   decimal.Decimal
}

type CreateProductRequest struct {
	Name             string          `json:"name" binding:"required,min=2,max=255"`
	Quantity         int64           `json:"quantity" binding:"required,min=1"`
	TotalCostOfGoods decimal.Decimal `json:"total_cost_of_goods" binding:"required,decimalgt=0"`
	TotalPriceSold   decimal.Decimal `json:"total_price_sold" binding:"required,decimalgt=0"`
}

type UpdateProductRequest struct {
	Name             string          `json:"name" binding:"required,min=2,max=255"`
	Quantity         int64           `json:"quantity" binding:"required,min=1"`
	TotalCostOfGoods decimal.Decimal `json:"total_cost_of_goods" binding:"required,decimalgt=0"`
	TotalPriceSold   decimal.Decimal `json:"total_price_sold" binding:"required,decimalgt=0"`
}
