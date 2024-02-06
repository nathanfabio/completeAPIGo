package validation

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/nathanfabio/completeAPIGo/internal/handler/httperr"
)

func ValidateHttpData(d interface{}) *httperr.RestErr {
	val := validator.New(validator.WithRequiredStructEnabled())

	//extract name
	val.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := val.Struct(d); err != nil {
		var errorsCauses []httperr.Fields

		for _, e := range err.(validator.ValidationErrors) {
			cause := httperr.Fields{}
			fieldName := e.Field()

			switch e.Tag() {
			case "required":
				cause.Message = fmt.Sprintf("%s is required", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "uuid4":
				cause.Message = fmt.Sprintf("%s is not valid uuid", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "boolean":
				cause.Message = fmt.Sprintf("%s is not valid boolean", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "min":
				cause.Message = fmt.Sprintf("%s must be greater than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "max":
				cause.Message = fmt.Sprintf("%s must be less than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "email":
				cause.Message = fmt.Sprintf("%s is not valid email", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "containsany":
				cause.Message = fmt.Sprintf("%s must contain at least one of these caracters: !@#$*", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			default:
				cause.Message = "invalid field"
				cause.Field = fieldName
				cause.Value = e.Value()
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return httperr.NewBadRequestValidationError("some fields are invalid", errorsCauses)
	}

	return nil
}
