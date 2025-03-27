package api

import (
	"github.com/Dupup1/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrencies validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}