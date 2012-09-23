package forms

type (
	Value struct {
		In string
		Go interface{}
	}

	valueMap map[string]Value

	errorMap map[string]error

	Result struct {
		Values valueMap
		Errors errorMap
	}
)
