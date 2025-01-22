package validationutils

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func DecimalGT(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.GreaterThan(baseValue)
}

func DecimalLT(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.LessThan(baseValue)
}

func DecimalGTE(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.GreaterThanOrEqual(baseValue)
}

func DecimalLTE(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	value, err := decimal.NewFromString(data)
	if err != nil {
		return false
	}

	baseValue, err := decimal.NewFromString(fl.Param())
	if err != nil {
		return false
	}
	return value.LessThanOrEqual(baseValue)
}

func TimeFormatValidator(fl validator.FieldLevel) bool {
	format := fl.Param()
	_, err := time.Parse(format, fl.Field().String())
	return err == nil
}
