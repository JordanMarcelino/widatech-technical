package repository

import (
	"context"
	"errors"

	"github.com/JordanMarcelino/widatech-technical/internal/dto"
	"github.com/JordanMarcelino/widatech-technical/internal/entity"
	"github.com/jackc/pgx/v5/pgconn"
)

type InvoiceRepository interface {
	Search(ctx context.Context, searchInvoiceRequest *dto.SearchInvoiceRequest) ([]*entity.Invoice, error)
	Save(ctx context.Context, invoice *entity.Invoice) error
	Update(ctx context.Context, invoice *entity.Invoice) error
	DeleteByID(ctx context.Context, invoiceNo string) error
}

type invoiceRepositoryImpl struct {
	DB DBTX
}

func NewInvoiceRepository(db DBTX) InvoiceRepository {
	return &invoiceRepositoryImpl{
		DB: db,
	}
}

func (r *invoiceRepositoryImpl) Search(ctx context.Context, searchInvoiceRequest *dto.SearchInvoiceRequest) ([]*entity.Invoice, error) {
	panic("implement me")
}

func (r *invoiceRepositoryImpl) Save(ctx context.Context, invoice *entity.Invoice) error {
	query := `
		insert into invoices (invoice_no, invoice_date, customer_name, sales_person_name, payment_type, notes)
		values ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.ExecContext(ctx, query, invoice.InvoiceNo, invoice.InvoiceDate, invoice.CustomerName, invoice.SalesPersonName, invoice.PaymentType, invoice.Notes)
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		return errors.New("invoice number already exists")
	}

	return err
}

func (r *invoiceRepositoryImpl) Update(ctx context.Context, invoice *entity.Invoice) error {
	panic("implement me")
}

func (r *invoiceRepositoryImpl) DeleteByID(ctx context.Context, invoiceNo string) error {
	panic("implement me")
}
