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
				Name: name,
				Validators: Validators{
					PresenceValidator},
			},
		}}

	badInputs := []string{empty, space, tab}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != PresenceError {
			t.Fatalf("Expected %v. Got %v - in: %#v", PresenceError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v.In != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}

	goodInputs := []string{"a", "aa", "a b c"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v", e)
		}

		if v, _ := res.Values[name]; v.In != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestNumericValidator(t *testing.T) {
	const name = "age"

	frm := &Form{
		Fields{
			{
				Name: name,
				Validators: Validators{
					NumericValidator},
			},
		}}

	badInputs := []string{empty, space, tab, "a"}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, ok := res.Errors[name]; !ok || e != NumericError {
			t.Fatalf("Expected %v. Got %v - in: %#v", NumericError, e, in)
		}

		if v, ok := res.Values[name]; !ok || v.In != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}

	goodInputs := []string{"1", "11", "1.1", "111"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v", e)
		}

		if v, _ := res.Values[name]; v.In != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestMaxLengthValidator(t *testing.T) {
	const name = "tweet"

	frm := &Form{
		Fields{
			{
				Name: name,
				Validators: Validators{
					MaxLengthValidator(2)},
			},
		}}

	badInputs := []string{"   ", "123"}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		m := "is too long (maximum is 2 characters)"
		if e, ok := res.Errors[name]; !ok || e.Error() != m {
			t.Fatalf("Got %v - in: %v", e, in)
		}

		if v, ok := res.Values[name]; !ok || v.In != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}

	goodInputs := []string{empty, space, tab, "1", "22"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v - in: %v", e, in)
		}

		if v, _ := res.Values[name]; v.In != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestMinLengthValidator(t *testing.T) {
	const name = "password"

	frm := &Form{
		Fields{
			{
				Name: name,
				Validators: Validators{
					MinLengthValidator(2, true)},
			},
		}}

	badInputs := []string{space, tab, "1"}
	for _, in := range badInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		m := "is too short (minimum is 2 characters)"
		if e, ok := res.Errors[name]; !ok || e.Error() != m {
			t.Fatalf("Got %v - in: %v", e, in)
		}

		if v, ok := res.Values[name]; !ok || v.In != in {
			t.Fatalf("Expected %#v. Got %#v", in, v)
		}
	}

	goodInputs := []string{empty, "22", "333"}
	for _, in := range goodInputs {
		res := frm.Load(httpRequest(url.Values{name: {in}}))

		if e, _ := res.Errors[name]; e != nil {
			t.Fatalf("Got %#v - in: %v", e, in)
		}

		if v, _ := res.Values[name]; v.In != in {
			t.Fatalf("Expected %v. Got %v", in, v)
		}
	}
}

func TestConfirmationOfValidator(t *testing.T) {

	frm := &Form{
		Fields{{
			Name:       "password",
			Validators: Validators{ConfirmationValidator("password")}}, {
			Name:       "password_confirmation",
			Validators: Validators{}},
		}}

	// goodInputs
	res := frm.Load(httpRequest(url.Values{
		"password":              {"secret"},
		"password_confirmation": {"secret"},
	}))

	if e, _ := res.Errors["password"]; e != nil {
		t.Fatalf("Got %#v", e)
	}

	if v, _ := res.Values["password"]; v.In != "secret" {
		t.Fatalf("Got %v", v)
	}

	if e, _ := res.Errors["password_confirmation"]; e != nil {
		t.Fatalf("Got %#v", e)
	}

	if v, _ := res.Values["password_confirmation"]; v.In != "secret" {
		t.Fatalf("Got %v", v)
	}

	// badInputs
	res = frm.Load(httpRequest(url.Values{
		"password":              {"secret"},
		"password_confirmation": {"secrets"},
	}))

	if e, _ := res.Errors["password"]; e != ConfirmationError {
		t.Fatalf("Got %v", e)
	}

	if v, _ := res.Values["password"]; v.In != "secret" {
		t.Fatalf("Got %v", v)
	}

	if e, _ := res.Errors["password_confirmation"]; e != nil {
		t.Fatalf("Got %#v", e)
	}

	if v, _ := res.Values["password_confirmation"]; v.In != "secrets" {
		t.Fatalf("Got %v", v)
	}
}
