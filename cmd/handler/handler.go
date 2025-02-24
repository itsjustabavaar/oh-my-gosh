package handler

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/auth"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func InputHandler(input string) (string, *int, error) {
	var state utils.State
	switch {
	case strings.HasPrefix(input, "mynameis"):
		state = &auth.MyNameIsState{}
	case strings.HasPrefix(input, "logout"):
		state = &auth.LogOutState{}
	case strings.HasPrefix(input, "whoami"):
		state = &auth.WhoAmIState{}
	case strings.HasPrefix(input, "exit"):
		state = &basiccommands.ExitState{}
	case strings.HasPrefix(input, "echo"):
		state = &basiccommands.EchoState{}
	case strings.HasPrefix(input, "pwd"):
		state = &basiccommands.PwdState{}
	default:
		return input, nil, nil
	}
	return state.Handler(input)
}
