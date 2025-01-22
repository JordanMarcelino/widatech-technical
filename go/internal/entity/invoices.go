package entity

type Invoice struct {
	InvoiceNo       string
	InvoiceDate     string
	CustomerName    string
	SalesPersonName string
	PaymentType     string
	Notes           string
	Products        []*Product
}
