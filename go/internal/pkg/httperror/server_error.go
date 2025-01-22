package httperror

import (
	"errors"
	"net/http"

	"github.com/JordanMarcelino/widatech-technical/internal/pkg/constant"
)

func NewServerError() *ResponseError {
	msg := constant.InternalServerErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusInternalServerError, msg)
}
