package auth

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
)

const addUserCommand = "adduser"

type AddUserCommand struct {
	Input  string
	Output string
}

func (a *AddUserCommand) Handler() {
	components := utils.SplitInput(a.Input, " ")
	if len(components) == 1 {
		utils.PrintError(addUserCommand, utils.ErrUsernameNotEntered)
		return
	}
	if len(components) > 2 {
		utils.PrintError(addUserCommand, utils.ErrTooManyArguments)
		return
	}

	username := components[1]

	password, err := utils.PasswordReader("enter password")
	if err != nil {
		utils.PrintError(addUserCommand, err)
		return
	}

	passwordAgain, err := utils.PasswordReader("enter password again")
	if err != nil {
		utils.PrintError(addUserCommand, err)
		return
	}

	if password != passwordAgain {
		utils.PrintError(addUserCommand, utils.ErrMissMatchPasswords)
		return
	}

	err = user.AddUser(username, password)
	if err != nil {
		utils.PrintError(addUserCommand, err)
		return
	}

	a.Output = fmt.Sprintf("user %s created successfully", username)
	utils.PrintOutput(a.Output)
}
