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
		Fields{
			{
				Name:       name,
				Validators: Validators{PresenceValidator},
			},
		}}

	badInputs := []string{empty, space, tab}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != PresenceError {
			t.Fatalf("Expected %v. Got %v - in: %#v", PresenceError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}

	goodInputs := []string{"a", "aa", "a b c"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v", e)
		}

		if v, _ := res.Values[name]; v != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestNumericValidator(t *testing.T) {
	const name = "age"

	frm := &Form{
		Fields{
			{
				Name:       name,
				Validators: Validators{NumericValidator},
			},
		}}

	badInputs := []string{empty, space, tab, "a"}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != NumericError {
			t.Fatalf("Expected %v. Got %v - in: %#v", NumericError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}

	goodInputs := []string{"1", "11", "1.1", "111"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v", e)
		}

		if v, _ := res.Values[name]; v != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestMaxLengthValidator(t *testing.T) {
	const name = "tweet"

	frm := &Form{
		Fields{
			{
				Name:       name,
				Validators: Validators{MaxLengthValidator(2)},
			},
		}}

	badInputs := []string{"   ", "123"}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		m := "is too long (maximum is 2 characters)"
		if e, ok := res.Errors[name]; !ok || e.Error() != m {
			t.Fatalf("Got %v - in: %v", e, in)
		}

		if v, ok := res.Values[name]; !ok || v != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}

	goodInputs := []string{empty, space, tab, "1", "22"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v - in: %v", e, in)
		}

		if v, _ := res.Values[name]; v != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}
