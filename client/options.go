package client

type ErrorOption func(*RequestError)

func WithPath(path string) ErrorOption {
	return func(ere *RequestError) {
		ere.Path = path
	}
}

func WithMethod(method string) ErrorOption {
	return func(ere *RequestError) {
		ere.Method = method
	}
}

func WithParams(param string) ErrorOption {
	return func(ere *RequestError) {
		ere.Param = param
	}
}

func WithCode(code string) ErrorOption {
	return func(ere *RequestError) {
		ere.Code = code
	}
}

func WithMessage(message string) ErrorOption {
	return func(ere *RequestError) {
		ere.Message = message
	}
}

func WithBody(body []byte) ErrorOption {
	return func(ere *RequestError) {
		ere.Body = body
	}
}

func WithError(err error) ErrorOption {
	return func(ere *RequestError) {
		ere.Err = err
	}
}
