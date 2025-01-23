package entity

import "github.com/shopspring/decimal"

type Invoice struct {
	InvoiceNo       string
	InvoiceDate     string
	CustomerName    string
	SalesPersonName string
	PaymentType     string
	Notes           string
	Products        []*Product
}

type InvoiceTransaction struct {
	TotalProfit decimal.Decimal
	TotalCash   decimal.Decimal
}
