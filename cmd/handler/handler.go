package handler

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/auth"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func InputHandler(input string) (string, *int, error) {
	var state utils.Command
	switch {
	case strings.HasPrefix(input, "login"):
		state = &auth.LoginCommand{Input: input}
	case strings.HasPrefix(input, "logout"):
		state = &auth.LogOutCommand{}
	case strings.HasPrefix(input, "whoami"):
		state = &auth.WhoAmICommand{}
	case strings.HasPrefix(input, "exit"):
		state = &basiccommands.ExitCommand{Input: input}
	case strings.HasPrefix(input, "echo"):
		state = &basiccommands.EchoCommand{Input: input}
	case strings.HasPrefix(input, "pwd"):
		state = &basiccommands.PwdCommand{}
	default:
		return input, nil, nil
	}
	return state.Handler()
}
