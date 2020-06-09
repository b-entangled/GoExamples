package errors

import (
	"fmt"
)

// Error :- Define Error Interface and Implement error interface (Error()string)
type Error interface {
	error
	Code() int
}

// CustomError :- Custom Error Fields
type CustomError struct {
	Message string
	ECode   int
}

// Initialize Error Interface
func CError(message string, code int) Error {
	return &CustomError{Message: message, ECode: code}
}

// Implementation of error Interface
func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %d", ce.Message, ce.ECode)
}

// Custom Function for Error Code
func (ce *CustomError) Code() int {
	return ce.ECode
}
