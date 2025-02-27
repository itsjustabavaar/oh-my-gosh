package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

const loginCommand = "login"

type LoginCommand struct {
	Input  string
	Output string
}

func (l *LoginCommand) Execute() {
	components := utils.SplitInput(l.Input, " ")
	if len(components) > 2 {
		utils.PrintError(loginCommand, utils.ErrTooManyArguments)
		return
	}

	username := strings.TrimSpace(strings.TrimPrefix(l.Input, "login "))
	if username == "" || l.Input == "login" {
		utils.PrintError(loginCommand, utils.ErrWhoYouAre)
		return
	}

	password, err := utils.PasswordReader("enter password")
	if err != nil {
		utils.PrintError(loginCommand, err)
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
