package validation

import (
	"log"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	ja_locales "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

func RegisterJaTranslation() ut.Translator {
	ja := ja_locales.New()
	jaUt := ut.New(ja, ja)
	translator, _ := jaUt.GetTranslator(ja.Locale())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			if jaName := field.Tag.Get("ja"); len(jaName) > 0 {
				return jaName
			}
			return ""
		})
		if err := ja_translations.RegisterDefaultTranslations(v, translator); err != nil {
			log.Fatalf("Failed to register default ja translation")
		}
	} else {
		log.Fatalf("Failed to register default ja translation")
	}
	return translator
}
