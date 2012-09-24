package forms

import "net/http"

type (
	Fields []Field

	Form struct {
		Fields Fields
	}
)

func (frm *Form) Load(req *http.Request) *Result {
	vals := make(valueMap)
	errs := make(errorMap)

	// copy each input, in its own loop so multi-field validators can
	// see across more than one field
	for _, fld := range frm.Fields {
		vals[fld.Name] = Value{
			In: req.FormValue(fld.Name),
		}
	}

	// validate each input
nextfield:
	for _, fld := range frm.Fields {

		err := fld.Validate(vals[fld.Name], vals)
		if err != nil {
			errs[fld.Name] = err
			continue nextfield
		}

		// TODO convert valid input
	}

	return &Result{Values: vals, Errors: errs}
}
