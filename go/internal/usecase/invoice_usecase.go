package usecase

import (
	"context"

	"github.com/JordanMarcelino/widatech-technical/internal/constant"
	"github.com/JordanMarcelino/widatech-technical/internal/dto"
	"github.com/JordanMarcelino/widatech-technical/internal/entity"
	"github.com/JordanMarcelino/widatech-technical/internal/httperror"
	phttperror "github.com/JordanMarcelino/widatech-technical/internal/pkg/httperror"
	"github.com/JordanMarcelino/widatech-technical/internal/pkg/utils/pageutils"
	"github.com/JordanMarcelino/widatech-technical/internal/repository"
)

type InvoiceUseCase interface {
	Search(ctx context.Context, req *dto.SearchInvoiceRequest) (*dto.SearchInvoiceResponse, *dto.PageMetaData, error)
	Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error)
	Update(ctx context.Context, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error)
	Delete(ctx context.Context, req *dto.DeleteInvoiceRequest) error
}

type invoiceUseCaseImpl struct {
	dataStore         repository.DataStore
	invoiceRepository repository.InvoiceRepository
}

func NewInvoiceUseCase(dataStore repository.DataStore, invoiceRepository repository.InvoiceRepository) InvoiceUseCase {
	return &invoiceUseCaseImpl{
		dataStore:         dataStore,
		invoiceRepository: invoiceRepository,
	}
}

func (u *invoiceUseCaseImpl) Search(ctx context.Context, req *dto.SearchInvoiceRequest) (*dto.SearchInvoiceResponse, *dto.PageMetaData, error) {
	invoices, err := u.invoiceRepository.Search(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	tx, err := u.invoiceRepository.CountTransaction(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	items, metaData := pageutils.CreateMetaData(invoices, req.Page, req.Limit)
	res := &dto.SearchInvoiceResponse{
		Invoices:    dto.ToInvoiceResponses(items),
		TotalCash:   tx.TotalCash,
		TotalProfit: tx.TotalProfit,
	}

	return res, metaData, nil
}

func (u *invoiceUseCaseImpl) Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	res := new(dto.InvoiceResponse)
	err := u.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
		invoiceRepository := ds.InvoiceRepository()
		productRepository := ds.ProductRepository()

		if ok := invoiceRepository.IsExistsByID(ctx, req.InvoiceNo); ok {
			return httperror.NewInvoiceDuplicateError(req.InvoiceNo)
		}

		invoice := &entity.Invoice{
			InvoiceNo:       req.InvoiceNo,
			InvoiceDate:     req.InvoiceDate,
			CustomerName:    req.CustomerName,
			SalesPersonName: req.SalesPersonName,
			PaymentType:     req.PaymentType,
			Notes:           req.Notes,
		}
		if err := invoiceRepository.Save(ctx, invoice); err != nil {
			return err
		}

		products := []*entity.Product{}
		for _, product := range req.Products {
			products = append(products, &entity.Product{
				InvoiceNo:        invoice.InvoiceNo,
				Name:             product.Name,
				Quantity:         product.Quantity,
				TotalCostOfGoods: product.TotalCostOfGoods,
				TotalPriceSold:   product.TotalPriceSold,
			})
		}

		if err := productRepository.SaveBatch(ctx, products); err != nil {
			return err
		}

		invoice.Products = products
		res = dto.ToInvoiceResponse(invoice)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *invoiceUseCaseImpl) Update(ctx context.Context, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error) {
	res := new(dto.InvoiceResponse)
	err := u.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
		invoiceRepository := ds.InvoiceRepository()
		productRepository := ds.ProductRepository()

		if ok := invoiceRepository.IsExistsByID(ctx, req.InvoiceNo); !ok {
			return phttperror.NewNotFoundError(constant.InvoiceNotFoundErrorMessage)
		}

		invoice := &entity.Invoice{
			InvoiceNo:       req.InvoiceNo,
			InvoiceDate:     req.InvoiceDate,
			CustomerName:    req.CustomerName,
			SalesPersonName: req.SalesPersonName,
			PaymentType:     req.PaymentType,
			Notes:           req.Notes,
		}
		if err := invoiceRepository.Update(ctx, invoice); err != nil {
			return err
		}

		if err := productRepository.DeleteByInvoiceNo(ctx, req.InvoiceNo); err != nil {
			return err
		}

		products := []*entity.Product{}
		for _, product := range req.Products {
			products = append(products, &entity.Product{
				InvoiceNo:        invoice.InvoiceNo,
				Name:             product.Name,
				Quantity:         product.Quantity,
				TotalCostOfGoods: product.TotalCostOfGoods,
				TotalPriceSold:   product.TotalPriceSold,
			})
		}

		if err := productRepository.SaveBatch(ctx, products); err != nil {
			return err
		}

		invoice.Products = products
		res = dto.ToInvoiceResponse(invoice)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *invoiceUseCaseImpl) Delete(ctx context.Context, req *dto.DeleteInvoiceRequest) error {
	return u.dataStore.Atomic(ctx, func(ds repository.DataStore) error {
		invoiceRepository := ds.InvoiceRepository()
		productRepository := ds.ProductRepository()

		if ok := invoiceRepository.IsExistsByID(ctx, req.InvoiceNo); !ok {
			return phttperror.NewNotFoundError(constant.InvoiceNotFoundErrorMessage)
		}

		if err := productRepository.DeleteByInvoiceNo(ctx, req.InvoiceNo); err != nil {
			return err
		}

		return invoiceRepository.DeleteByID(ctx, req.InvoiceNo)
	})
}
