package bloomfilter

type InvalidArgumentError struct {
	m    string
	name string
	v    interface{}
}

func NewInvalidArgumentError(m, name string, v interface{}) InvalidArgumentError {
	return InvalidArgumentError{m, name, v}
}

func (err InvalidArgumentError) Error() string {
	return err.m
}
