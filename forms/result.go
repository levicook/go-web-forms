package forms

type valueMap map[string]string
type errorMap map[string]error

type Result struct {
	Values valueMap
	Errors errorMap
}
