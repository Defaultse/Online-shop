package translator

import (
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/kk"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
)

var (
	TransRU ut.Translator
	TransEN ut.Translator
	TransKZ ut.Translator
)

func init() {
	english := en.New()
	russian := ru.New()
	kazakh := kk.New()

	trans := ut.New(russian, russian, kazakh, english)

	err := trans.Import(ut.FormatJSON, "./translations/response")
	if err != nil {
		log.Fatal(err)
	}

	err = trans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

	TransRU, _ = trans.GetTranslator(russian.Locale())
	TransEN, _ = trans.GetTranslator(english.Locale())
	TransKZ, _ = trans.GetTranslator(kazakh.Locale())
}

func Translate(message string, lang string) string {
	switch lang {
	case "ru":
		translation, _ := TransRU.T(message)
		return translation
	case "en":
		translation, _ := TransEN.T(message)
		return translation
	case "kz":
		translation, _ := TransKZ.T(message)
		return translation
	default:
		translation, _ := TransRU.T(message)
		return translation
	}
}
