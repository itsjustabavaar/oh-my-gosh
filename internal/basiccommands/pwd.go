package basiccommands

import (
	"os"
)

type PwdCommand struct {
	Output string
}

func (p *PwdCommand) Handler() (string, *int, error) {
	var err error
	p.Output, err = os.Getwd()
	return p.Output, nil, err
}
