package provider

import (
	"github.com/JordanMarcelino/widatech-technical/internal/config"
	"github.com/JordanMarcelino/widatech-technical/internal/controller"
	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {
	appController := controller.NewAppController()
	appController.Route(router)

	BootstrapInvoice(router)
}
