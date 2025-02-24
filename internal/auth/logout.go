package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

type LogOutState struct {
	Output string
}

func (l *LogOutState) Handler(input string) (string, *int, error) {
	var err error
	components := strings.Split(input, " ")
	if len(components) > 1 {
		err = utils.ErrTooManyArguments
	}
	if vars.Prompt == "$ " {
		err = utils.ErrAlreadyLoggedOut
	} else {
		vars.Prompt = "$ "
	}
	return l.Output, nil, err
}
