package basiccommands

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
)

type GoshCommand struct {
	Output string
}

func (g *GoshCommand) Handler() {
	g.Output = "I'm Here! :)"
	utils.PrintOutput(g.Output)
}
