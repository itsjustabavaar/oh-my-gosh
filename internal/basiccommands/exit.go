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
	var err error
	var code *int = nil
	components := strings.Split(e.Input, " ")
	if len(components) > 2 {
		err = utils.ErrTooManyArguments
	} else if len(components) == 1 {
		exit := 0
		code = &exit
	} else {
		exit, convErr := strconv.Atoi(components[1])
		if convErr != nil {
			err = utils.ErrInvalidExitCode
		} else {
			code = &exit
			if vars.Prompt != "$ " {
				e.Output = fmt.Sprintf("%s logged out", strings.TrimSuffix(vars.Prompt, "$ "))
			}
		}
	}
	return e.Output, code, err
}
