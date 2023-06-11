package oasaSyncValidation

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

const erroMessageTemplate = "Field validation for [%s] failed on the [%s] tag"

type OpswValidateError struct {
	Key     string
	Message string
}

var OpswValidation *validator.Validate

func InitializeValidation() {
	OpswValidation = validator.New()
}

func OpswValidate(input interface{}) bool {
	var isValid = true
	err := OpswValidation.Struct(input)
	if err != nil {
		analyzeError(err, nil)
		isValid = false
	}
	return isValid
}

func analyzeError(err error, trans ut.Translator) (errs []OpswValidateError) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errs = append(errs, OpswValidateError{
			Key:     e.Field(),
			Message: fmt.Sprintf(erroMessageTemplate, e.Field(), e.Tag()),
		})
	}
	return errs
}
