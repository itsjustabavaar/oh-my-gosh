package exit

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
)

const exitCommand = "exit"

type ExitCommand struct {
	Input  string
	Output string
}

func (e *ExitCommand) Execute() {
	var exitCode *int = nil

	components := strings.Split(e.Input, " ")
	if len(components) > 2 {
		util.PrintError(exitCommand, errorutil.ErrTooManyArguments)
		return
	}

	if len(components) == 1 {
		code := 0
		exitCode = &code
	} else {
		code, err := strconv.Atoi(components[1])
		if err != nil {
			util.PrintError(exitCommand, errorutil.ErrInvalidExitCode)
			return
		}

		exitCode = &code
	}

	if exitCode != nil {
		e.Output = fmt.Sprintf("exit status %d", *exitCode)
		util.PrintOutput(e.Output)
		os.Exit(*exitCode)
	}
}
