package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

type WhoAmICommand struct {
	Output string
}

func (w *WhoAmICommand) Handler() (string, *int, error) {
	var err error
	if vars.Prompt != "$ " {
		w.Output = strings.TrimSuffix(vars.Prompt, "$ ")
	} else {
		w.Output = "Anonymous"
	}
	return w.Output, nil, err
}
