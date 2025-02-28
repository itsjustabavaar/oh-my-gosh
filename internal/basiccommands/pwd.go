package basiccommands

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
)

const pwdCommand = "pwdCommand"

type PwdCommand struct {
	Output string
}

func (p *PwdCommand) Execute() {
	p.Output, _ = os.Getwd()
	utils.PrintOutput(p.Output)
}
