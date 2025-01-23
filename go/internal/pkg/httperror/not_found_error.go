package httperror

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/JordanMarcelino/widatech-technical/internal/pkg/constant"
)

func NewNotFoundError(msg string) *ResponseError {
	msg = fmt.Sprintf(constant.NotFoundErrorMessage, msg)
	err := errors.New(msg)

	return NewResponseError(err, http.StatusNotFound, msg)
}
