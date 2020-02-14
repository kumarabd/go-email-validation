package errors

type (
	Error struct {
		Code        string
		Description string
	}
)

func New(code string, des string, doc ...string) *Error {
	return &Error{
		Code:        code,
		Description: des,
	}
}

func (e *Error) Error() string { return e.Description }

func Is(err error) (*Error, bool) {
	if err != nil {
		e, ok := err.(*Error)
		return e, ok
	}
	return nil, false
}
