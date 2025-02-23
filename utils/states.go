package utils

type State interface {
	Handler(input string, prompt *string) (string, error)
}
