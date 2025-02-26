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
		utils.Error(exitCommand, utils.ErrTooManyArguments)
		return
	}

	if len(components) == 1 {
		code := 0
		exitCode = &code
	} else {
		code, err := strconv.Atoi(components[1])
		if err != nil {
			utils.Error(exitCommand, utils.ErrInvalidExitCode)
			return
		}
		exitCode = &code
	}

	if vars.CurrentUser.Username != "" {
		_, _ = fmt.Fprintf(vars.StandardOutput, "%s logged out\n", vars.CurrentUser.Username)
	}

	if exitCode != nil {
		os.Exit(*exitCode)
	}
}
