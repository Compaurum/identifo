package mem

//Error - domain level error type
type Error string

//Error - implementation of std.Error protocol
func (e Error) Error() string { return string(e) }

const (
	//ErrorNotFound general not found error
	ErrorNotFound = Error("not found")
)
