package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type GoshCommand struct {
	Output string
}

func (g *GoshCommand) Handler() {
	g.Output = "I'm Here! :)"
	_, _ = fmt.Fprintln(vars.StandardOutput, g.Output)
}
