package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

const loginCommand = "login"

type LoginCommand struct {
	Input  string
	Output string
}

func (l *LoginCommand) Execute() {
	components := utils.SplitInput(l.Input, " ")
	var username, password string

	switch len(components) {
	case 1:
		utils.PrintError(loginCommand, utils.ErrWhoYouAre)
		return
	case 2:
		username = components[1]
	case 3:
		username, password = components[1], components[2]
	default:
		utils.PrintError(loginCommand, utils.ErrTooManyArguments)
		return
	}

	loginUser, err := user.Login(username, password)
	if err != nil {
		utils.PrintError(loginCommand, err)
		return
	}

	vars.CurrentUser = loginUser

	l.Output = "login succeed"

	utils.PrintOutput(l.Output)
}
