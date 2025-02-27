package basiccommands

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
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
		utils.PrintError(pwdCommand, err)
		return
	}
	utils.PrintOutput(p.Output)
}
