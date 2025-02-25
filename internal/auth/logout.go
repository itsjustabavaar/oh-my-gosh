package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type LogOutCommand struct {
	Output string
}

func (l *LogOutCommand) Handler() (string, *int, error) {
	if vars.CurrentUser.Username == "" {
		return l.Output, nil, utils.ErrAlreadyLoggedOut
	}

	vars.CurrentUser = &user.User{}
	l.Output = "logout succeed"

	return l.Output, nil, nil
}
