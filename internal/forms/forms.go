package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f Form) has(field string, r *http.Request) bool {
	fieldExist := r.FormValue(field)

	return fieldExist != ""
}
