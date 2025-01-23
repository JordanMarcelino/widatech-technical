package controller

import (
	"github.com/JordanMarcelino/widatech-technical/internal/dto"
	"github.com/JordanMarcelino/widatech-technical/internal/pkg/utils/ginutils"
	"github.com/JordanMarcelino/widatech-technical/internal/pkg/utils/pageutils"
	"github.com/JordanMarcelino/widatech-technical/internal/usecase"
	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	invoiceUseCase usecase.InvoiceUseCase
}

func NewInvoiceController(invoiceUseCase usecase.InvoiceUseCase) *InvoiceController {
	return &InvoiceController{invoiceUseCase: invoiceUseCase}
}

func (c *InvoiceController) Route(r *gin.Engine) {
	r.GET("/invoices", c.Search)
	r.POST("/invoices", c.Create)
	r.PUT("/invoices/:invoice_no", c.Update)
	r.DELETE("/invoices/:invoice_no", c.Delete)
}

func (c *InvoiceController) Search(ctx *gin.Context) {
	req := new(dto.SearchInvoiceRequest)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.Error(err)
		return
	}

	res, paging, err := c.invoiceUseCase.Search(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	paging.Links = pageutils.CreateLinks(ctx.Request, int(paging.Page), int(paging.Size), int(paging.TotalItem), int(paging.TotalPage))
	ginutils.ResponseOKPagination(ctx, res, paging)
}

func (c *InvoiceController) Create(ctx *gin.Context) {
	req := new(dto.CreateInvoiceRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.invoiceUseCase.Create(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseCreated(ctx, res)
}

func (c *InvoiceController) Update(ctx *gin.Context) {

}

func (c *InvoiceController) Delete(ctx *gin.Context) {

}
