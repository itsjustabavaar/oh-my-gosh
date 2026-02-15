package gosh

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
)

type GoshCommand struct {
	Output string
}

func (g *GoshCommand) Execute() {
	g.Output = "I'm Here! :)"
	util.PrintOutput(g.Output)
}
