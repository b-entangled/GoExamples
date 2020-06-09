package errors

import(
	"fmt"
)

type Error interface{
	error
	Code()int
}

type CustomError struct{
	Message string
	ECode int
}

func CError(message string, code int) Error{
	return &CustomError{Message: message, ECode: code}
}

func (ce *CustomError) Error() string{
	return fmt.Sprintf("Error: %s, Code: %d", ce.Message, ce.ECode)
}
func (ce *CustomError) Code() int{
	return ce.ECode
}