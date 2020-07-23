package vo

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidFilePath return invalid FilePath
	ErrInvalidFilePath = errors.New("invalid file path")

	rxFilePath = regexp.MustCompile(`^\/[\w+-]*[\/\w+-]*\/?$`)
)

// FilePath structure
type FilePath struct {
	value string
}

// NewFilePath create new FilePath
func NewFilePath(value string) (FilePath, error) {
	var f = FilePath{value: value}

	if !f.validate() {
		return FilePath{}, ErrInvalidFilePath
	}

	return f, nil
}

func (f FilePath) validate() bool {
	return rxFilePath.MatchString(f.value)
}

// Value return value FilePath
func (f FilePath) Value() string {
	return f.value
}

// String returns string representation of the FilePath
func (f FilePath) String() string {
	return string(f.value)
}

// Equals checks that two FilePath are the same
func (f FilePath) Equals(value Value) bool {
	o, ok := value.(FilePath)
	return ok && f.value == o.value
}
