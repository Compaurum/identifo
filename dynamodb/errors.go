package dynamodb

//Error - domain level error type
type Error string

//Error - implementation of std.Error protocol
func (e Error) Error() string { return string(e) }

const (
	//ErrorInternalError internal error
	ErrorInternalError = Error("Internal error")
	// ErrorInactiveUser means user is inactive
	ErrorInactiveUser = Error("User is inactive")
)
