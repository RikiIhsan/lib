package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	ens      = en.New()
	uni      = ut.New(ens, ens)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
)

func Init(translator ...Translator) {
	en_translations.RegisterDefaultTranslations(validate, trans)
	for _, item := range translator {
		validate.RegisterTranslation(item.Tag, trans, func(ut ut.Translator) error {
			return ut.Add(item.Tag, item.Message, true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(item.Tag, fe.Field())
			return t
		})
	}
}

func Validate(data interface{}) []ValidatorError {
	validationErrors := []ValidatorError{}
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidatorError
			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Message = err.Translate(trans)
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
