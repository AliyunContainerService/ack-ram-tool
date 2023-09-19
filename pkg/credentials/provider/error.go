package provider

type NotEnableError struct {
	err error
}

func NewNotEnableError(err error) *NotEnableError {
	return &NotEnableError{err: err}
}
func (e NotEnableError) Error() string {
	return e.err.Error()
}

func isNotEnableError(err error) bool {
	if _, ok := err.(*NotEnableError); ok {
		return true
	}
	if _, ok := err.(NotEnableError); ok {
		return true
	}
	return false
}
