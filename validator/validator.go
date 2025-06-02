package validator

import (
	"reflect"
	"strings"

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

func GetFieldTag(data interface{}, fieldName string, sourceTag string) string {
	t := reflect.TypeOf(data)
	if t == nil {
		return fieldName
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fieldName
	}
	field, found := t.FieldByName(fieldName)
	if !found {
		return fieldName
	}
	tagValue := field.Tag.Get(sourceTag)
	if tagValue == "" || tagValue == "-" {
		return fieldName
	}
	return strings.Split(tagValue, ",")[0]
}

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

func Validate(data interface{}, source string) []ValidatorError {
	validationErrors := []ValidatorError{}
	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidatorError
			elem.FailedField = GetFieldTag(data, err.Field(), source)
			elem.Tag = err.Tag()
			elem.Message = err.Translate(trans)
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}

func SliceValidate(data interface{}, source string) []ValidatorError {
	validationErrors := []ValidatorError{}
	errs := validate.Var(data, "dive")
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ValidatorError
			elem.FailedField = GetFieldTag(data, err.Field(), source)
			elem.Tag = err.Tag()
			elem.Message = err.Translate(trans)
			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
