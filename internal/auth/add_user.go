package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
)

type AddUserCommand struct {
	Input  string
	Output string
}

func (a *AddUserCommand) Handler() (string, *int, error) {
	components := utils.SplitInput(a.Input, " ")
	if len(components) == 1 {
		return a.Output, nil, utils.ErrUsernameNotEntered
	}
	if len(components) != 2 {
		return a.Output, nil, utils.ErrInvalidArguments
	}

	username := components[1]

	password, err := utils.PasswordReader("enter password")
	if err != nil {
		return a.Output, nil, err
	}

	passwordAgain, err := utils.PasswordReader("enter password again")
	if err != nil {
		return a.Output, nil, err
	}

	if password != passwordAgain {
		return a.Output, nil, utils.ErrMissMatchPasswords
	}

	err = user.AddUser(username, password)
	if err != nil {
		return a.Output, nil, err
	}

	a.Output = "user created successfully"
	return a.Output, nil, nil
}
