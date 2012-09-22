package forms

import (
	"errors"
	"strings"
)

var (
	PresenceValidator ValidatorFunc = presenceValidator
	PresenceError     error         = errors.New("must be present")

	NumericValidator ValidatorFunc = numericValidator
	NumericError     error         = errors.New("must be numeric")
)

type Validator interface {
	Validate(in string) (out string, err error)
}

type ValidatorFunc func(string) (string, error)

func (f ValidatorFunc) Validate(in string) (out string, err error) {
	return f(in)
}

func presenceValidator(in string) (string, error) {
	if len(in) == 0 {
		return in, PresenceError
	}
	if len(strings.TrimSpace(in)) == 0 {
		return in, PresenceError
	}
	return in, nil
}

func numericValidator(in string) (out string, err error) {
	out, err = in, NumericError
	return
}

func MaxLengthValidator(min int) ValidatorFunc {
	return func(in string) (out string, err error) {
		return
	}
}

func MinLengthValidator(min int) ValidatorFunc {
	return func(in string) (out string, err error) {
		return
	}
}
