package provider

import (
	"github.com/JordanMarcelino/widatech-technical/internal/controller"
	"github.com/JordanMarcelino/widatech-technical/internal/repository"
	"github.com/JordanMarcelino/widatech-technical/internal/usecase"
	"github.com/gin-gonic/gin"
)

func BootstrapInvoice(router *gin.Engine) {
	invoiceRepository := repository.NewInvoiceRepository(db)

	invoiceUseCase := usecase.NewInvoiceUseCase(dataStore, invoiceRepository)
	invoiceController := controller.NewInvoiceController(invoiceUseCase)

	invoiceController.Route(router)
}
