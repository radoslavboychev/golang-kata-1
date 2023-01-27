package liberror

const (
	ErrFailedToOpenFile       = Error("failed to open file!")
	ErrFailedToResolveAuthors = Error("failed to resolve authors!")
	ErrFailedToPrint          = Error("failed to print product")
	ErrGeneric                = Error("failed to execute action")
	ErrFailedToFindProduct    = Error("failed to find product")
	ErrFailedToLoadConfig     = Error("failed to load configuration")
	ErrISBNInvalid            = Error("ISBN is of invalid length")
	ErrInvalidEmail           = Error("invalid email address")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
