package validator

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func validationRequired(v *validator.Validate) {
	translationFn := func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	}

	v.RegisterTranslation("required", transRU, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("required", transEN, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("required", transKZ, func(ut ut.Translator) error {
		return nil
	}, translationFn)
}

func validationPhoneNumber(v *validator.Validate) {
	translationFn := func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	}

	v.RegisterTranslation("phone", transRU, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("phone", transEN, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("phone", transKZ, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		phoneNumber := fl.Field().String()

		ok, _ := regexp.Match(`^[7]\d{9}$`, []byte(phoneNumber))

		return ok
	})
}

func validationConfirmationCode(v *validator.Validate) {
	translationFn := func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("confirmationCode", fe.Field())
		return t
	}

	v.RegisterTranslation("confirmationCode", transRU, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("confirmationCode", transEN, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterTranslation("confirmationCode", transKZ, func(ut ut.Translator) error {
		return nil
	}, translationFn)

	v.RegisterValidation("confirmationCode", func(fl validator.FieldLevel) bool {
		confirmationCode := fl.Field().String()

		return len(confirmationCode) == 8
	})

}
