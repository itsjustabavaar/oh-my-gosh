package utils

type State interface {
	Handler(input string) (string, *int, error)
}
