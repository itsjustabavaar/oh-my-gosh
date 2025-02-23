package handler

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/auth"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func InputHandler(input string, prompt *string) (output string, err error) {
	var state utils.State
	switch {
	case strings.HasPrefix(input, "mynameis "):
		state = &auth.MyNameIsState{}
	case input == "whoami":
		state = &auth.WhoAmIState{}
	case strings.HasPrefix(input, "exit "):
		state = &basiccommands.ExitState{}
	default:
		return input, nil
	}
	return state.Handler(input, prompt)
}
