package errors

type Error interface {
	Error() string
	Code() ErrorCode
}

type CommonError struct {
	code ErrorCode
	err  error
}

func NewCommonError(code ErrorCode) Error {
	return CommonError{
		code: code,
		err:  nil,
	}
}

func (ce CommonError) Code() ErrorCode {
	return ce.code
}

func (ce CommonError) Error() string {
	return ce.err.Error()
}
