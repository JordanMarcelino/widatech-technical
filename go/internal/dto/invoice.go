package dto

type InvoiceResponse struct {
	InvoiceNo       string
	InvoiceDate     string
	CustomerName    string
	SalesPersonName string
	PaymentType     string
	Notes           string
	Products        []ProductResponse
}

type CreateInvoiceRequest struct {
	InvoiceNo       string                 `json:"invoice_no" binding:"required,min=1,max=255"`
	InvoiceDate     string                 `json:"invoice_date" binding:"required,time_format=2006-01-02"`
	CustomerName    string                 `json:"customer_name" binding:"required,min=2,max=255"`
	SalesPersonName string                 `json:"sales_person_name" binding:"required,min=2,max=255"`
	PaymentType     string                 `json:"payment_type" binding:"required,oneof=CASH CREDIT"`
	Notes           string                 `json:"notes" binding:"required,min=5"`
	Products        []CreateProductRequest `json:"products" binding:"required,dive"`
}

type UpdateInvoiceRequest struct {
	InvoiceNo       string                 `json:"-"`
	InvoiceDate     string                 `json:"invoice_date" binding:"required,time_format=2006-01-02"`
	CustomerName    string                 `json:"customer_name" binding:"required,min=2,max=255"`
	SalesPersonName string                 `json:"sales_person_name" binding:"required,min=2,max=255"`
	PaymentType     string                 `json:"payment_type" binding:"required,oneof=CASH CREDIT"`
	Notes           string                 `json:"notes" binding:"required,min=5"`
	Products        []UpdateProductRequest `json:"products" binding:"required,dive"`
}

type SearchInvoiceRequest struct {
	StartDate string `form:"start-date" binding:"required,time_format=2006-01-02,ltefield=EndDate"`
	EndDate   string `form:"end-date" binding:"required,time_format=2006-01-02,gtefield=StartDate"`
	Page      int64  `form:"page" binding:"required,min=1"`
	Limit     int64  `form:"limit" binding:"required,min=1,max=25"`
}

type DeleteInvoiceRequest struct {
	InvoiceNo string `json:"-"`
}
