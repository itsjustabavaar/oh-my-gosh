package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

type LoginCommand struct {
	Input  string
	Output string
}

func (l *LoginCommand) Handler() (string, *int, error) {
	components := utils.SplitInput(l.Input, " ")
	if len(components) > 2 {
		return l.Output, nil, utils.ErrTooManyArguments
	}

	username := strings.TrimSpace(strings.TrimPrefix(l.Input, "login "))
	if username == "" || l.Input == "login" {
		return l.Output, nil, utils.ErrWhoYouAre
	}

	password, err := utils.PasswordReader("enter password")
	if err != nil {
		return l.Output, nil, err
	}

	loginUser, err := user.Login(username, password)
	if err != nil {
		return l.Output, nil, err
	}

	vars.CurrentUser = loginUser
	l.Output = "login succeed"

	return l.Output, nil, err
}
