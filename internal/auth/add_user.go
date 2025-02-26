package auth

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

const addUserCommand = "adduser"

type AddUserCommand struct {
	Input  string
	Output string
}

func (a *AddUserCommand) Handler() {
	components := utils.SplitInput(a.Input, " ")
	if len(components) == 1 {
		utils.Error(addUserCommand, utils.ErrUsernameNotEntered)
		return
	}
	if len(components) > 2 {
		utils.Error(addUserCommand, utils.ErrTooManyArguments)
		return
	}

	username := components[1]

	password, err := utils.PasswordReader("enter password")
	if err != nil {
		utils.Error(addUserCommand, err)
		return
	}

	passwordAgain, err := utils.PasswordReader("enter password again")
	if err != nil {
		utils.Error(addUserCommand, err)
		return
	}

	if password != passwordAgain {
		utils.Error(addUserCommand, utils.ErrMissMatchPasswords)
		return
	}

	err = user.AddUser(username, password)
	if err != nil {
		utils.Error(addUserCommand, err)
		return
	}

	a.Output = fmt.Sprintf("user %s create successfully", username)
	_, _ = fmt.Fprintln(vars.StandardOutput, a.Output)
}
