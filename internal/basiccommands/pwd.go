package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
)

const pwdCommand = "pwdCommand"

type PwdCommand struct {
	Output string
}

func (p *PwdCommand) Handler() {
	var err error
	p.Output, err = os.Getwd()
	if err != nil {
		utils.Error(pwdCommand, err)
		return
	}
	_, _ = fmt.Fprintln(vars.StandardOutput, p.Output)
}
