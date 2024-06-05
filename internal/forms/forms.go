package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}

// returns true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		x := f.Get(field)
		if strings.TrimSpace(x) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, d int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < d {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", d))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
