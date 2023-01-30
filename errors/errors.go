package liberror

const (
	ErrFailedToOpenFile             = Error("failed to open file! file is invalid or can not be found")
	ErrFailedToResolveAuthors       = Error("failed to resolve authors!")
	ErrFailedToPrint                = Error("failed to print product")
	ErrGeneric                      = Error("failed to execute action")
	ErrFailedToFindProduct          = Error("failed to find product")
	ErrFailedToLoadConfig           = Error("failed to load configuration")
	ErrISBNInvalid                  = Error("ISBN is of invalid length")
	ErrInvalidEmail                 = Error("invalid email address")
	ErrEmailIsNull                  = Error("email should not be empty")
	ErrNoProductsFoundWithAuthor    = Error("no products found with author")
	ErrFileNotFound                 = Error("file not found")
	ErrFilenameInvalid              = Error("filename invalid or null")
	ErrNoBooksLoaded                = Error("no books loaded from file!")
	ErrNoMagazinesLoaded            = Error("no magazines loaded")
	FailedToResolveAuthorsInvalid   = Error("failed to resolve authors, authors are invalid")
	FailedToResolveBooksInvalid     = Error("failed to resolve authors, books are invalid")
	FailedToResolveMagazinesInvalid = Error("failed to resolve authors, magazines are invalid")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
