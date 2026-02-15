package pwd

import (
	"os"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
)

const pwdCommand = "pwdCommand"

type PwdCommand struct {
	Output string
}

func (p *PwdCommand) Execute() {
	p.Output, _ = os.Getwd()
	util.PrintOutput(p.Output)
}
