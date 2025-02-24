package basiccommands

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
	"strings"
)

type PwdState struct {
	Output string
}

func (p *PwdState) Handler(input string) (string, *int, error) {
	var err error
	components := strings.Split(input, " ")
	if len(components) > 1 {
		err = utils.ErrTooManyArguments
	} else {
		p.Output, err = os.Getwd()
	}
	return p.Output, nil, err
}
