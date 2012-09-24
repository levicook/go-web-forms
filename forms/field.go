package forms

type (
	Validators []Validator

	Field struct {
		Name       string
		Validators Validators
		Converter  Converter
	}
)

func (fld *Field) Validate(val Value, vals valueMap) (err error) {

	for _, v := range fld.Validators {
		err = v.Validate(val, vals)
		if err != nil {
			return
		}
	}

	return
}
