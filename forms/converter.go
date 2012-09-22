package forms

type Converter interface {
	Convert(in string) (out interface{}, err error)
}
