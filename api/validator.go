package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/zaenalarifin12/simplebank/utils"
)

var validatorCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {

	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
