package login

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

const loginCommand = "login"

type LoginCommand struct {
	Input  string
	Output string
}

func (l *LoginCommand) Execute() {
	components := util.SplitInput(l.Input, " ")
	var username, password string

	switch len(components) {
	case 1:
		util.PrintError(loginCommand, errorutil.ErrWhoYouAre)
		return
	case 2:
		username = components[1]
	case 3:
		username, password = components[1], components[2]
	default:
		util.PrintError(loginCommand, errorutil.ErrTooManyArguments)
		return
	}

	loginUser, err := user.Login(username, password)
	if err != nil {
		util.PrintError(loginCommand, err)
		return
	}

	vars.CurrentUser = loginUser

	l.Output = "login succeed"

	util.PrintOutput(l.Output)
}
