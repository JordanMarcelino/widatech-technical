package entity

import "github.com/shopspring/decimal"

type Product struct {
	ID               int64
	InvoiceNo        string
	Name             string
	Quantity         int64
	TotalCostOfGoods decimal.Decimal
	TotalPriceSold   decimal.Decimal
}
