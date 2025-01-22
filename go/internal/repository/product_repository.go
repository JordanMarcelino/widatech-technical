package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/JordanMarcelino/widatech-technical/internal/entity"
)

type ProductRepository interface {
	SaveBatch(ctx context.Context, products []*entity.Product) error
	DeleteByInvoiceNo(ctx context.Context, invoiceNo string) error
}

type productRepositoryImpl struct {
	DB DBTX
}

func NewProductRepository(db DBTX) ProductRepository {
	return &productRepositoryImpl{
		DB: db,
	}
}

func (r *productRepositoryImpl) SaveBatch(ctx context.Context, products []*entity.Product) error {
	params := []string{}
	args := []any{}

	for i, product := range products {
		params = append(params, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", i*5+1, i*5+2, i*5+3, i*5+4, i*5+5))
		args = append(args, product.InvoiceNo)
		args = append(args, product.Name)
		args = append(args, product.Quantity)
		args = append(args, product.TotalCostOfGoods)
		args = append(args, product.TotalPriceSold)
	}

	query := fmt.Sprintf("insert into products (invoice_no, name, quantity, total_cost_of_goods, total_price_sold) values %s returning id", strings.Join(params, ","))

	rows, err := r.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if err := rows.Scan(&products[i].ID); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *productRepositoryImpl) DeleteByInvoiceNo(ctx context.Context, invoiceNo string) error {
	query := `
		delete from products where invoice_no = $1
	`

	_, err := r.DB.ExecContext(ctx, query, invoiceNo)
	return err
}
