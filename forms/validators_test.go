package forms

import (
	"net/url"
	"testing"
)

const (
	empty = ""
	space = " "
	tab   = "	"
)

func TestPresenceValidator(t *testing.T) {
	const name = "username"

	frm := &Form{
		Fields: Fields{
			{
				Name:       name,
				Validators: Validators{PresenceValidator},
			},
		}}

	inputs := []string{empty, space, tab}
	for _, in := range inputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != PresenceError {
			t.Fatalf("Expected %v. Got %v - in: %#v", PresenceError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestNumericValidator(t *testing.T) {
	const name = "age"

	frm := &Form{
		Fields: Fields{
			{
				Name:       name,
				Validators: Validators{NumericValidator},
			},
		}}

	inputs := []string{empty, space, tab, "a"}
	for _, in := range inputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != NumericError {
			t.Fatalf("Expected %v. Got %v - in: %#v", NumericError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}
}
