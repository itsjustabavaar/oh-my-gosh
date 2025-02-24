package utils

type Command interface {
	Handler() (string, *int, error)
}
