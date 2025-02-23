package utils

type HandlerState uint8

const (
	StateUnknown HandlerState = iota
	StateMyNameIs
	StateWhoAmI
	StateExit
)

type State interface {
	Apply()
}

type WhoAmIState struct {
}
