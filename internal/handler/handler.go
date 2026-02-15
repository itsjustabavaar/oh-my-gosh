package handler

import (
	"strings"

	"github.com/itsjustabavaar/oh-my-gosh/internal/commands"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/auth/adduser"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/auth/login"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/auth/logout"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/auth/whoami"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/cat"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/cd"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/commandtype"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/echo"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/exit"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/gosh"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/history"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/home"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/pwd"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/system"
	"github.com/itsjustabavaar/oh-my-gosh/internal/redirection"
)

func InputHandler(input string) {
	var command commands.Command

	file := redirection.DetectAndApplyRedirection(&input)

	switch {
	case strings.HasPrefix(input, "adduser"):
		command = &adduser.AddUserCommand{Input: input}
	case strings.HasPrefix(input, "login"):
		command = &login.LoginCommand{Input: input}
	case strings.HasPrefix(input, "logout"):
		command = &logout.LogOutCommand{}
	case strings.HasPrefix(input, "whoami"):
		command = &whoami.WhoAmICommand{}
	case strings.HasPrefix(input, "exit"):
		command = &exit.ExitCommand{Input: input}
	case strings.HasPrefix(input, "echo"):
		command = &echo.EchoCommand{Input: input}
	case strings.HasPrefix(input, "pwd"):
		command = &pwd.PwdCommand{}
	case strings.HasPrefix(input, "cat"):
		command = &cat.CatCommand{Input: input}
	case strings.HasPrefix(input, "cd"):
		command = &cd.CdCommand{Input: input}
	case strings.HasPrefix(input, "~"):
		command = &home.HomeCommand{}
	case strings.HasPrefix(input, "gosh"):
		command = &gosh.GoshCommand{}
	case strings.HasPrefix(input, "history"):
		command = &history.HistoryCommand{Input: input}
	case strings.HasPrefix(input, "type"):
		command = &commandtype.TypeCommand{Input: input}
	default:
		command = &system.SystemCommand{Input: input}
	}

	command.Execute()

	redirection.RestoreStates(file)
}
