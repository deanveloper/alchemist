package alchemist

import "strconv"

// SyntaxError represents an error made while
// trying to parse an Alchemist file.
type SyntaxError struct {
	reason string
	line   int
}

func (s *SyntaxError) Error() string {
	str := s.reason
	if s.line > 0 {
		str += " on line " + strconv.Itoa(s.line)
	}
	return str
}

// UnknownError occurs when an error was caused by some
// other outside source, such as an I/O error.
type UnknownError struct {
	reason string
	cause  error
}

func (u *UnknownError) Error() string {
	return u.reason + " caused by " + u.cause.Error()
}

// Unwrap implements the golang.org/x/xerrors Wrapper interface.
func (u *UnknownError) Unwrap() error {
	return u.cause
}
