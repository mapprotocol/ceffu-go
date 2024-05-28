package client

import "fmt"

type RequestError struct {
	Path    string
	Method  string
	Param   string
	Code    string
	Message string
	Body    []byte
	Err     error
}

func NewRequestError(path string, opts ...ErrorOption) *RequestError {
	ext := RequestError{
		Path: path,
	}
	for _, o := range opts {
		o(&ext)
	}
	return &ext
}

func (e *RequestError) Error() string {
	msg := fmt.Sprintf("Error while making request to %s,", e.Path)

	if e.Method != "" {
		msg += fmt.Sprintf(" method: %s, ", e.Method)
	}
	if e.Param != "" {
		msg += fmt.Sprintf(" param: %s, ", e.Param)
	}
	if e.Code != "" {
		msg += fmt.Sprintf("code: %s, ", e.Code)
	}
	if e.Message != "" {
		msg += fmt.Sprintf("message: %s, ", e.Message)
	}
	if len(e.Body) > 0 {
		msg += fmt.Sprintf("body: %s, ", string(e.Body))
	}
	if e.Err != nil {
		msg += fmt.Sprintf("error: %s", e.Err.Error())
	}
	return msg
}
