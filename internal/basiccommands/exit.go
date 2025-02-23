package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strconv"
	"strings"
)

type ExitState struct {
	Output   string
	ExitCode int
}

func (e *ExitState) Handler(input string, prompt *string) (output string, err error) {
	components := strings.Split(input, " ")
	if len(components) > 2 {
		err = utils.ErrTooManyArguments
	}
	exit, convErr := strconv.Atoi(components[1])
	if convErr != nil {
		err = utils.ErrInvalidExitCode
	}
	e.ExitCode = exit
	if *prompt != "$ " {
		e.Output = fmt.Sprintf("%s logged out\nexit status %d", strings.TrimPrefix(*prompt, "$ "), exit)
	} else {
		e.Output = fmt.Sprintf("exit status %d", exit)
	}
	return e.Output, err
}
