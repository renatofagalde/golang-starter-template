package validation

import (
	"encoding/json"
	"errors"
	"github.com/HunCoding/golang-architecture/hexagonal-news-api/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"time"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateError(
	validation_err error,
) *rest_err.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	var timeParseErr *time.ParseError

	errorsCauses := []rest_err.Causes{}

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses = errorsAppend(validation_err, errorsCauses)
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else if errors.As(validation_err, &timeParseErr) {
		cause := rest_err.Causes{
			Message: timeParseErr.Message,
			Field:   timeParseErr.Value,
		}
		errorsCauses = append(errorsCauses, cause)
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}

func errorsAppend(validation_err error, errorsCauses []rest_err.Causes) []rest_err.Causes {
	for _, e := range validation_err.(validator.ValidationErrors) {
		cause := rest_err.Causes{
			Message: e.Translate(transl),
			Field:   e.Field(),
		}

		errorsCauses = append(errorsCauses, cause)
	}
	return errorsCauses
}
