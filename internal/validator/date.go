package validator

import (
	"time"

	playvalidator "github.com/go-playground/validator/v10"
)

func notFuture(fl playvalidator.FieldLevel) bool {
	switch v := fl.Field().Interface().(type) {
	case time.Time:
		return !v.After(time.Now())
	case *time.Time:
		if v == nil {
			return true
		}
		return !v.After(time.Now())
	default:
		return false
	}
}
