package validator

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/kk"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"

	"github.com/go-playground/validator/v10"
	en_translator "github.com/go-playground/validator/v10/translations/en"
	ru_translator "github.com/go-playground/validator/v10/translations/ru"
)

var (
	transRU ut.Translator
	transEN ut.Translator
	transKZ ut.Translator
	trans   *ut.UniversalTranslator
)

func MakeValidator() *validator.Validate {
	english := en.New()
	russian := ru.New()
	kazakh := kk.New()

	trans = ut.New(russian, russian, kazakh, english)

	transEN, _ = trans.FindTranslator("en", english.Locale())
	transRU, _ = trans.FindTranslator("ru", russian.Locale())
	transKZ, _ = trans.FindTranslator("kk", kazakh.Locale())

	err := trans.Import(ut.FormatJSON, "./translations/validator")
	if err != nil {
		log.Fatal(err)
	}

	err = trans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

	v := validator.New()

	ru_translator.RegisterDefaultTranslations(v, transRU)
	en_translator.RegisterDefaultTranslations(v, transEN)

	// Registration of field custom validators and translators
	lowerCaseTag(v)
	validationRequired(v)
	validationPhoneNumber(v)
	validationConfirmationCode(v)

	return v
}

func lowerCaseTag(v *validator.Validate) {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

func translateError(valErr validator.FieldError, lang string) string {
	switch lang {
	case "ru":
		return valErr.Translate(transRU)
	case "en":
		return valErr.Translate(transEN)
	case "kz":
		return valErr.Translate(transKZ)
	default:
		return valErr.Translate(transRU)
	}
}
