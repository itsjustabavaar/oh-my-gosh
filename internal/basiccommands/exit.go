package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strconv"
	"strings"
)

type ExitCommand struct {
	Input  string
	Output string
}

func (e *ExitCommand) Handler() (string, *int, error) {
	var exitCode *int = nil

	components := strings.Split(e.Input, " ")
	if len(components) > 2 {
		return e.Output, nil, utils.ErrTooManyArguments
	}

	if len(components) == 1 {
		code := 0
		exitCode = &code
	} else {
		code, err := strconv.Atoi(components[1])
		if err != nil {
			return e.Output, nil, utils.ErrInvalidExitCode
		} else {
			exitCode = &code
		}
	}

	if vars.CurrentUser.Username != "" {
		e.Output = fmt.Sprintf("%s logged out", vars.CurrentUser.Username)
	}

	return e.Output, exitCode, nil
}
