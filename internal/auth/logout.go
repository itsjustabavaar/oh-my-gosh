package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type LogOutCommand struct {
	Output string
}

func (l *LogOutCommand) Handler() (string, *int, error) {
	var err error
	if vars.Prompt == "$ " {
		err = utils.ErrAlreadyLoggedOut
	} else {
		vars.Prompt = "$ "
		l.Output = "logout succeed"
	}
	return l.Output, nil, err
}
