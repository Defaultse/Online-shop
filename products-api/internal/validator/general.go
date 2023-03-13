package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator interface {
}

func ValidateCommon[T Validator](dto T, lang string) (ok bool, errors map[string]interface{}) {
	v := MakeValidator()
	ok = true
	errors = make(map[string]interface{})
	err := v.Struct(dto)

	if err != nil {
		valErrs := err.(validator.ValidationErrors)

		for _, valErr := range valErrs {
			errors[valErr.Field()] = translateError(valErr, lang)
		}

		ok = false
	}

	return ok, errors
}
