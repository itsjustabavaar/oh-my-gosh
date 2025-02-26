package systemcommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"os/exec"
)

type SystemCommand struct {
	Input  string
	Output string
}

func (s *SystemCommand) Handler() {
	components := utils.SplitInput(s.Input, " ")
	command, arguments := components[0], components[1:]
	_, err := basiccommands.FindingType(command)
	if err == nil {
		cmd := exec.Command(command, arguments...)
		cmd.Dir, _ = os.Getwd()
		output, err := cmd.CombinedOutput()
		if err != nil {
			_, _ = fmt.Fprintln(vars.StandardError, err)
			return
		}

		_, _ = fmt.Fprint(vars.StandardOutput, string(output))
		return
	}
	_, _ = fmt.Fprintln(vars.StandardError, err)
}
