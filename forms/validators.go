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

type (
	Validator interface {
		Validate(val Value, vals valueMap) error
	}

	ValidatorFunc func(Value, valueMap) error
)

func (f ValidatorFunc) Validate(val Value, vals valueMap) error {
	return f(val, vals)
}

func presenceValidator(val Value, vals valueMap) error {

	if len(val.In) == 0 {
		return PresenceError
	}

	if len(strings.TrimSpace(val.In)) == 0 {
		return PresenceError
	}

	return nil
}

func numericValidator(val Value, vals valueMap) error {

	if _, e := strconv.ParseFloat(val.In, 64); e != nil {
		return NumericError
	}

	return nil
}

// TODO InclusionValidator // in a set
// TODO ExclusionValidator // not in a set

// TODO func TimeValidator
// TODO func MinTimeValidator
// TODO func MaxTimeValidator

// TODO func DurationValidator

// TODO func BetweenLengthValidator(min max int, inclusive bool) ??

func MinLengthValidator(min int, emptyOk bool) ValidatorFunc {
	return func(val Value, vals valueMap) error {

		if emptyOk && val.In == "" {
			return nil
		}

		if len(val.In) < min {
			m := fmt.Sprintf("is too short (minimum is %v characters)", min)
			return errors.New(m)
		}

		return nil
	}
}

func MaxLengthValidator(max int) ValidatorFunc {
	return func(val Value, vals valueMap) error {

		if len(val.In) > max {
			m := fmt.Sprintf("is too long (maximum is %v characters)", max)
			return errors.New(m)
		}

		return nil
	}
}

func ConfirmationValidator(fldName string) ValidatorFunc {
	return func(val Value, vals valueMap) error {

		cf := fmt.Sprintf("%v_confirmation", fldName)

		if v, _ := vals[cf]; v.In != val.In {
			return ConfirmationError
		}

		return nil
	}
}
