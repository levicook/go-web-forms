package forms

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	PresenceValidator ValidatorFunc = presenceValidator
	PresenceError     error         = errors.New("must be present")

	NumericValidator ValidatorFunc = numericValidator
	NumericError     error         = errors.New("must be numeric")

	ConfirmationError error = errors.New("doesn't match confirmation")
)

type Validator interface {
	Validate(in string, vals valueMap) (out string, err error)
}

type ValidatorFunc func(string, valueMap) (string, error)

func (f ValidatorFunc) Validate(in string, vals valueMap) (out string, err error) {
	return f(in, vals)
}

// TODO Range based validations
//type Range interface {
//  Start() interface{}
//  End() interface{}
//  Include(interface{}) bool
//}

func presenceValidator(in string, vals valueMap) (string, error) {
	if len(in) == 0 {
		return in, PresenceError
	}
	if len(strings.TrimSpace(in)) == 0 {
		return in, PresenceError
	}
	return in, nil
}

func numericValidator(in string, vals valueMap) (string, error) {
	if _, e := strconv.ParseFloat(in, 64); e != nil {
		return in, NumericError
	}
	return in, nil
}

func MaxLengthValidator(max int) ValidatorFunc {
	return func(in string, vals valueMap) (string, error) {

		if len(in) > max {
			m := fmt.Sprintf("is too long (maximum is %v characters)", max)
			return in, errors.New(m)
		}

		return in, nil
	}
}

func MinLengthValidator(min int, emptyOk bool) ValidatorFunc {
	return func(in string, vals valueMap) (string, error) {

		if emptyOk && in == "" {
			return in, nil
		}

		if len(in) < min {
			m := fmt.Sprintf("is too short (minimum is %v characters)", min)
			return in, errors.New(m)
		}

		return in, nil
	}
}

func ConfirmationValidator(fldName string) ValidatorFunc {
	return func(in string, vals valueMap) (string, error) {

		cf := fmt.Sprintf("%v_confirmation", fldName)

		if v, _ := vals[cf]; v != in {
			return in, ConfirmationError
		}

		return in, nil
	}
}
