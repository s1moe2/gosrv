package handlers

import (
	"encoding/json"
	"net/http"
)

type customError interface {
	StatusCode() int
}

type ErrorList []error

func (el ErrorList) Error() string {
	ret := ""
	for _, err := range el {
		if len(ret) > 0 {
			ret = ret + ", "
		}

		ret = ret + err.Error()
	}
	return ret
}

func (el ErrorList) MarshalJSON() ([]byte, error) {
	data := []byte("[")
	for i, err := range el {
		if i != 0 {
			data = append(data, ',')
		}

		errBytes, err := json.Marshal(err.Error())
		if err != nil {
			return nil, err
		}

		data = append(data, errBytes...)
	}
	data = append(data, ']')

	return data, nil
}

type userError struct {
	Status int       `json:"status"`
	Errors ErrorList `json:"errors"`
}

// newSimpleUserError returns a new userError with a default Bad Request status code
// and the list of errors passed as argument
func newUserError(errs []error) *userError {
	return &userError{
		Status: http.StatusBadRequest,
		Errors: errs,
	}
}

// newSimpleUserError returns a new userError with a default Bad Request status code
// and creates a list of errors from the one error passed as argument
func newSimpleUserError(err error) *userError {
	return &userError{
		Status: http.StatusBadRequest,
		Errors: []error{err},
	}
}

func (ue userError) StatusCode() int {
	return ue.Status
}

type appError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (ae appError) StatusCode() int {
	return ae.Status
}

// internalError returns an appError with 500 code and default message
func internalError() appError {
	return appError{
		Status:  http.StatusInternalServerError,
		Message: "Internal server error",
	}
}
