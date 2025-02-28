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

func (a *AddUserCommand) Execute() {
	components := utils.SplitInput(a.Input, " ")
	var username, password string

	switch len(components) {
	case 1:
		utils.PrintError(addUserCommand, utils.ErrUsernameNotEntered)
		return
	case 2:
		username = components[1]
	case 3:
		username, password = components[1], components[2]
	default:
		utils.PrintError(addUserCommand, utils.ErrTooManyArguments)
		return
	}

	err := user.AddUser(username, password)
	if err != nil {
		utils.PrintError(addUserCommand, err)
		return
	}

	a.Output = fmt.Sprintf("user %s created successfully", username)
	utils.PrintOutput(a.Output)
}
