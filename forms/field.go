package forms

type Validators []Validator

type Field struct {
	Name       string
	Validators Validators
	Converter  Converter
}

func (fld *Field) Validate(in string) (out string, err error) {
	out = in

	for _, v := range fld.Validators {
		out, err = v.Validate(in)
		if err != nil {
			return
		}
	}

	return
}
