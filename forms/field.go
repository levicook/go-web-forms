package forms

type Validators []Validator

type Field struct {
	Name       string
	Validators Validators
	Converter  Converter
}

func (fld *Field) Validate(in string, vals valueMap) (out string, err error) {
	out = in

	for _, v := range fld.Validators {
		out, err = v.Validate(in, vals)
		if err != nil {
			return
		}
	}

	return
}
