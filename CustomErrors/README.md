Example to Implement and Use Custom Error Types.

errors.go :-

// Error :- Define Error Interface and Implement error interface (Error()string)
// Implements error interface along with custom methods
type Error interface {
	error
	Code() int
}


// CustomError :- Custom Error Fields
type CustomError struct {
	Message string
	ECode   int
}

// Implementation of error Interface
func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error: %s, Code: %d", ce.Message, ce.ECode)
}
