package adduser

import (
	"fmt"

	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
)

const addUserCommand = "adduser"

type AddUserCommand struct {
	Input  string
	Output string
}

func (a *AddUserCommand) Execute() {
	components := util.SplitInput(a.Input, " ")
	var username, password string

	switch len(components) {
	case 1:
		util.PrintError(addUserCommand, errorutil.ErrUsernameNotEntered)
		return
	case 2:
		username = components[1]
	case 3:
		username, password = components[1], components[2]
	default:
		util.PrintError(addUserCommand, errorutil.ErrTooManyArguments)
		return
	}

	err := user.AddUser(username, password)
	if err != nil {
		util.PrintError(addUserCommand, err)
		return
	}

	a.Output = fmt.Sprintf("user %s created successfully", username)
	util.PrintOutput(a.Output)
}
