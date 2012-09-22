package forms

import (
	"net/http"
	"net/url"
)

func httpRequest(values url.Values) (req *http.Request) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	req.Form = values
	return
}
