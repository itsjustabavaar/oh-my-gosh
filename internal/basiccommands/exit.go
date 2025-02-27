package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"strconv"
	"strings"
)

const exitCommand = "exit"

type ExitCommand struct {
	Input  string
	Output string
}

func (e *ExitCommand) Handler() {
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

	if vars.CurrentUser.Username != "" {
		e.Output = fmt.Sprintf("%s logged out\n", vars.CurrentUser.Username)
		utils.PrintOutput(e.Output)
	}

	if exitCode != nil {
		os.Exit(*exitCode)
	}
}
