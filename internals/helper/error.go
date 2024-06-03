package helper

import (
	"errors"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnIfError(err error) {
	if err != nil {
		panic(&httpError{err, http.StatusInternalServerError})
	}
}

type httpError struct {
	error
	statusCode int
}

func IsErrorNotFound(err error) bool {
	if err == nil {
		return false
	}

	var httpError *httpError
	if errors.As(err, &httpError) && httpError.statusCode == http.StatusNotFound {
		return true
	}
	return false
}