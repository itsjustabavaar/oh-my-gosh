package utils

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrDuplicateUser       = errors.New("duplicate user exists with this username")
	ErrCommandNotFound     = errors.New("command not found")
	ErrTooManyArguments    = errors.New("too many arguments")
	ErrUnknownType         = errors.New("unknown type")
	ErrEmptyCommandHistory = errors.New("empty command history")
	ErrInvalidUsername     = errors.New("invalid username")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrInvalidExitCode     = errors.New("invalid exit code")
	ErrAlreadyLoggedOut    = errors.New("already logged out")
	ErrNotEnoughArguments  = errors.New("not enough arguments")
)
