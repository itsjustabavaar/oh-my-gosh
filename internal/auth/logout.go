package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

type LogOutState struct {
	Output string
}

func (l *LogOutState) Handler(input string, prompt *string) (output string, err error) {
	components := strings.Split(input, " ")
	if len(components) > 1 {
		err = utils.ErrTooManyArguments
	}
	if *prompt == "$ " {
		err = utils.ErrInvalidUsername
	} else {
		*prompt = "$ "
	}
	return l.Output, err
}
