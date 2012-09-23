package forms

type Validators []Validator

type Field struct {
	Name       string
	Validators Validators
	Converter  Converter
}

func (fld *Field) Validate(in string, frm *Form) (out string, err error) {
	out = in

	for _, v := range fld.Validators {
		out, err = v.Validate(in, frm)
		if err != nil {
			return
		}
	}

	return
}
