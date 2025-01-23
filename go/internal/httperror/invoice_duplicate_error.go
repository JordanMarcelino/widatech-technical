package httperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JordanMarcelino/widatech-technical/internal/constant"
	"github.com/JordanMarcelino/widatech-technical/internal/pkg/httperror"
)

func NewInvoiceDuplicateError(invoiceNo string) *httperror.ResponseError {
	msg := fmt.Sprintf(constant.InvoiceDuplicateErrorMessage, invoiceNo)
	err := errors.New(msg)

	return httperror.NewResponseError(err, http.StatusBadRequest, msg)
}
