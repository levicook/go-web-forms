package forms

import "net/http"

type Fields []Field

type Form struct {
	Fields Fields
}

func (frm *Form) Load(req *http.Request) *Result {
	vals := make(valueMap)
	errs := make(errorMap)

nextfield:
	for _, fld := range frm.Fields {

		// copy the input
		vals[fld.Name] = req.FormValue(fld.Name)

		// validate the input
		_, err := fld.Validate(vals[fld.Name])
		if err != nil {
			errs[fld.Name] = err
			continue nextfield
		}

		// TODO convert the input

	}

	return &Result{Values: vals, Errors: errs}
}
