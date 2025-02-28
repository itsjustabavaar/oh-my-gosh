package systemcommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"os/exec"
	"runtime"
)

type SystemCommand struct {
	Input  string
	Output string
}

func (s *SystemCommand) Execute() {
	components := utils.SplitInput(s.Input, " ")

	command, arguments := components[0], components[1:]
	if runtime.GOOS == "windows" {
		_, ok := vars.WindowsBuiltins[command]
		if ok {
			command = "cmd"
			arguments = []string{"/C"}
			arguments = append(arguments, components...)
		}
	}

	_, err := basiccommands.FindingType(command)
	if err != nil {
		_, _ = fmt.Fprintln(vars.StandardError, err)
		return
	}
	cmd := exec.Command(command, arguments...)

	cmd.Dir, _ = os.Getwd()

	output, err := cmd.CombinedOutput()
	if err != nil {
		_, _ = fmt.Fprintln(vars.StandardError, err)
		return
	}

	utils.PrintOutput(string(output))
	return
}
