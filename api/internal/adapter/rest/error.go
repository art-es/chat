package rest

import "errors"

type httpError struct {
	error
	status int
}

func newErrorString(errmsg string, status int) *httpError {
	return &httpError{errors.New(errmsg), status}
}
