package app

type ErrorLevel int

const (
	ErrorLevelLow ErrorLevel = iota
	ErrorLevelMedium
	ErrorLevelHigh
)

type Error struct {
	message string
	level   ErrorLevel
}

func (err Error) Error() string {
	return err.message
}

func (err Error) Level() ErrorLevel {
	return err.level
}

func NewError(m string) Error {
	return Error{message: m, level: ErrorLevelLow}
}

func NewHighLevelError(m string) Error {
	return Error{message: m, level: ErrorLevelHigh}
}

var (
	ErrNotFound        = NewError("not found")
	ErrAccessDenied    = NewError("access denied")
	ErrInvalidUsername = NewError("invalid username")
	ErrInvalidPassword = NewError("invalid password")
)
