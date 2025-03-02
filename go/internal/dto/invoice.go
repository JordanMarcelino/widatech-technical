package dto

import (
	"github.com/JordanMarcelino/widatech-technical/internal/entity"
	"github.com/shopspring/decimal"
)

type SearchInvoiceResponse struct {
	Invoices    []*InvoiceResponse `json:"invoices"`
	TotalCash   decimal.Decimal    `json:"total_cash"`
	TotalProfit decimal.Decimal    `json:"total_profit"`
}

type InvoiceResponse struct {
	InvoiceNo       string             `json:"invoice_no"`
	InvoiceDate     string             `json:"invoice_date"`
	CustomerName    string             `json:"customer_name"`
	SalesPersonName string             `json:"sales_person_name"`
	PaymentType     string             `json:"payment_type"`
	Notes           string             `json:"notes"`
	Products        []*ProductResponse `json:"products"`
}

type CreateInvoiceRequest struct {
	InvoiceNo       string                  `json:"invoice_no" binding:"required,min=1,max=255"`
	InvoiceDate     string                  `json:"invoice_date" binding:"required,time_format=2006-01-02"`
	CustomerName    string                  `json:"customer_name" binding:"required,min=2,max=255"`
	SalesPersonName string                  `json:"sales_person_name" binding:"required,min=2,max=255"`
	PaymentType     string                  `json:"payment_type" binding:"required,oneof='CASH' 'CREDIT'"`
	Notes           string                  `json:"notes" binding:"required,min=5"`
	Products        []*CreateProductRequest `json:"products" binding:"required,min=1,dive"`
}

type UpdateInvoiceRequest struct {
	InvoiceNo       string                  `json:"-"`
	InvoiceDate     string                  `json:"invoice_date" binding:"required,time_format=2006-01-02"`
	CustomerName    string                  `json:"customer_name" binding:"required,min=2,max=255"`
	SalesPersonName string                  `json:"sales_person_name" binding:"required,min=2,max=255"`
	PaymentType     string                  `json:"payment_type" binding:"required,oneof=CASH CREDIT"`
	Notes           string                  `json:"notes" binding:"required,min=5"`
	Products        []*UpdateProductRequest `json:"products" binding:"required,dive"`
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

func ToInvoiceResponses(invoices []*entity.Invoice) []*InvoiceResponse {
	res := []*InvoiceResponse{}

	for _, invoice := range invoices {
		res = append(res, ToInvoiceResponse(invoice))
	}

	return res
}

func ToInvoiceResponse(invoice *entity.Invoice) *InvoiceResponse {
	return &InvoiceResponse{
		InvoiceNo:       invoice.InvoiceNo,
		InvoiceDate:     invoice.InvoiceDate,
		CustomerName:    invoice.CustomerName,
		SalesPersonName: invoice.SalesPersonName,
		PaymentType:     invoice.PaymentType,
		Notes:           invoice.Notes,
		Products:        ToProductResponses(invoice.Products),
	}
}
