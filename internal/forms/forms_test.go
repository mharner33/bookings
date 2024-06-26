package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData

	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	if form.Has("whatever") {
		t.Error("form shows has field when it does not")
	}

	postData := url.Values{}
	postData.Add("a", "a")

	form = New(postData)
	if !form.Has("a") {
		t.Error("form shows does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	postData.Add("some_field", "123")

	r := httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData

	form := New(r.PostForm)
	if !form.MinLength("some_field", 3) {
		t.Error("shows min length does not work when it should")
	}

	if form.MinLength("some_field", 4) {
		t.Error("shows min length does work when it should not")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}
	postData := url.Values{}
	postData.Add("email", "mike@test.com")

	form = New(postData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}
}
