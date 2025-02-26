package handler

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/auth"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/internal/systemcommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func InputHandler(input string) {
	var command utils.Command

	switch {
	case strings.HasPrefix(input, "adduser"):
		command = &auth.AddUserCommand{Input: input}
	case strings.HasPrefix(input, "login"):
		command = &auth.LoginCommand{Input: input}
	case strings.HasPrefix(input, "logout"):
		command = &auth.LogOutCommand{}
	case strings.HasPrefix(input, "whoami"):
		command = &auth.WhoAmICommand{}
	case strings.HasPrefix(input, "exit"):
		command = &basiccommands.ExitCommand{Input: input}
	case strings.HasPrefix(input, "echo"):
		command = &basiccommands.EchoCommand{Input: input}
	case strings.HasPrefix(input, "pwd"):
		command = &basiccommands.PwdCommand{}
	case strings.HasPrefix(input, "cat"):
		command = &basiccommands.CatCommand{Input: input}
	case strings.HasPrefix(input, "cd"):
		command = &basiccommands.CdCommand{Input: input}
	case strings.HasPrefix(input, "~"):
		command = &basiccommands.HomeCommand{}
	case strings.HasPrefix(input, "gosh"):
		command = &basiccommands.GoshCommand{}
	case strings.HasPrefix(input, "history"):
		command = &basiccommands.HistoryCommand{}
	case strings.HasPrefix(input, "type"):
		command = &basiccommands.TypeCommand{Input: input}
	default:
		//_, _ = fmt.Fprintln(vars.StandardOutput, input)
		command = &systemcommands.SystemCommand{Input: input}
	}

	command.Handler()
}
