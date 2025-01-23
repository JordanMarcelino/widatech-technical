package repository

import (
	"context"
	"errors"

	"github.com/JordanMarcelino/widatech-technical/internal/dto"
	"github.com/JordanMarcelino/widatech-technical/internal/entity"
	"github.com/jackc/pgx/v5/pgconn"
)

type InvoiceRepository interface {
	Search(ctx context.Context, req *dto.SearchInvoiceRequest) ([]*entity.Invoice, error)
	Save(ctx context.Context, invoice *entity.Invoice) error
	Update(ctx context.Context, invoice *entity.Invoice) error
	DeleteByID(ctx context.Context, invoiceNo string) error
	IsExistsByID(ctx context.Context, invoiceNo string) bool
	CountTransaction(ctx context.Context, req *dto.SearchInvoiceRequest) (*entity.InvoiceTransaction, error)
}

type invoiceRepositoryImpl struct {
	DB DBTX
}

func NewInvoiceRepository(db DBTX) InvoiceRepository {
	return &invoiceRepositoryImpl{
		DB: db,
	}
}

func (r *invoiceRepositoryImpl) Search(ctx context.Context, req *dto.SearchInvoiceRequest) ([]*entity.Invoice, error) {
	query := `
		select
			i.invoice_no, i.invoice_date, i.customer_name, i.sales_person_name, i.payment_type, i.notes,
			p.id, p.name, p.quantity, p.total_cost_of_goods, p.total_price_sold
		from
			invoices i
		join
			products p on i.invoice_no = p.invoice_no
		where
			i.invoice_date between $1 and $2
	`

	rows, err := r.DB.QueryContext(ctx, query, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoiceMap := map[string]*entity.Invoice{}
	for rows.Next() {
		invoice := new(entity.Invoice)
		product := new(entity.Product)

		if err := rows.Scan(
			&invoice.InvoiceNo,
			&invoice.InvoiceDate,
			&invoice.CustomerName,
			&invoice.SalesPersonName,
			&invoice.PaymentType,
			&invoice.Notes,
			&product.ID,
			&product.Name,
			&product.Quantity,
			&product.TotalCostOfGoods,
			&product.TotalPriceSold,
		); err != nil {
			return nil, err
		}

		product.InvoiceNo = invoice.InvoiceNo
		if _, ok := invoiceMap[invoice.InvoiceNo]; !ok {
			invoice.Products = []*entity.Product{}
			invoiceMap[invoice.InvoiceNo] = invoice
		}

		invoiceMap[invoice.InvoiceNo].Products = append(invoiceMap[invoice.InvoiceNo].Products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	invoices := []*entity.Invoice{}
	for _, invoice := range invoiceMap {
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

func (r *invoiceRepositoryImpl) Save(ctx context.Context, invoice *entity.Invoice) error {
	query := `
		insert into invoices
			(invoice_no, invoice_date, customer_name, sales_person_name, payment_type, notes)
		values
			($1, $2, $3, $4, $5, $6)
	`

	_, err := r.DB.ExecContext(ctx, query, invoice.InvoiceNo, invoice.InvoiceDate, invoice.CustomerName, invoice.SalesPersonName, invoice.PaymentType, invoice.Notes)
	if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
		return errors.New("invoice number already exists")
	}

	return err
}

func (r *invoiceRepositoryImpl) Update(ctx context.Context, invoice *entity.Invoice) error {
	query := `
		update invoices set invoice_date = $2, customer_name = $3, sales_person_name = $4, payment_type = $5, notes = $6
		where invoice_no = $1
	`

	_, err := r.DB.ExecContext(ctx, query, invoice.InvoiceNo, invoice.InvoiceDate, invoice.CustomerName, invoice.SalesPersonName, invoice.PaymentType, invoice.Notes)

	return err
}

func (r *invoiceRepositoryImpl) DeleteByID(ctx context.Context, invoiceNo string) error {
	query := `
		delete from invoices where invoice_no = $1
	`

	_, err := r.DB.ExecContext(ctx, query, invoiceNo)
	return err
}

func (r *invoiceRepositoryImpl) IsExistsByID(ctx context.Context, invoiceNo string) bool {
	query := `
	select exists(
		select * from invoices where invoice_no = $1
		)
		`

	var exists bool
	if err := r.DB.QueryRowContext(ctx, query, invoiceNo).Scan(&exists); err != nil {
		return false
	}

	return exists
}

func (r *invoiceRepositoryImpl) CountTransaction(ctx context.Context, req *dto.SearchInvoiceRequest) (*entity.InvoiceTransaction, error) {
	query := `
		with total_profit as(
			select
				sum(p.total_price_sold - p.total_cost_of_goods) total_profit
			from
				invoices i
			join
				products p on i.invoice_no = p.invoice_no
			where
				i.invoice_date between $1 and $2
		),
		total_cash as (
    		select
				sum(p.total_price_sold) total_cash
			from
				invoices i
			join
				products p on i.invoice_no = p.invoice_no
			where
				i.invoice_date between $1 and $2 and i.payment_type = 'CASH'
		)
		select
			tp.total_profit, tc.total_cash
		from
			total_cash tc, total_profit tp
	`

	tx := new(entity.InvoiceTransaction)
	if err := r.DB.QueryRowContext(ctx, query, req.StartDate, req.EndDate).Scan(&tx.TotalProfit, &tx.TotalCash); err != nil {
		return nil, err
	}

	return tx, nil
}
