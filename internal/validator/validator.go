package validator

import playvalidator "github.com/go-playground/validator/v10"

type Validator struct {
	validate *playvalidator.Validate
}

func New() *Validator {
	v := playvalidator.New()

	v.RegisterValidation("notfuture", notFuture)

	return &Validator{
		validate: v,
	}
}

func (v *Validator) Struct(s any) error {
	return v.validate.Struct(s)
}
