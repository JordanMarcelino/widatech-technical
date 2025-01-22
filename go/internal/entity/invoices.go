package entity

import "time"

type Invoice struct {
	InvoiceNo       string
	InvoiceDate     time.Time
	CustomerName    string
	SalesPersonName string
	PaymentType     string
	Notes           string
	Products        []Product
}
