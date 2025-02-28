package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
	"strconv"
	"strings"
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
		utils.PrintError(exitCommand, utils.ErrTooManyArguments)
		return
	}

	if len(components) == 1 {
		code := 0
		exitCode = &code
	} else {
		code, err := strconv.Atoi(components[1])
		if err != nil {
			utils.PrintError(exitCommand, utils.ErrInvalidExitCode)
			return
		}

		exitCode = &code
	}

	if exitCode != nil {
		e.Output = fmt.Sprintf("exit status %d", *exitCode)
		utils.PrintOutput(e.Output)
		os.Exit(*exitCode)
	}
}
