package system

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/commandtype"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

type SystemCommand struct {
	Input  string
	Output string
}

func (s *SystemCommand) Execute() {
	components := util.SplitInput(s.Input, " ")

	command, arguments := components[0], components[1:]
	if runtime.GOOS == "windows" {
		_, ok := vars.WindowsBuiltins[command]
		if ok {
			command = "cmd"
			arguments = []string{"/C"}
			arguments = append(arguments, components...)
		}
	}

	_, err := commandtype.FindingType(command)
	if err != nil {
		_, _ = fmt.Fprintln(vars.StandardError, err)
		return
	}
	cmd := exec.Command(command, arguments...)

	cmd.Dir, _ = os.Getwd()

	output, _ := cmd.CombinedOutput()

	util.PrintOutput(string(output))
}
