package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

type WhoAmIState struct {
	Output string
}

func (w *WhoAmIState) Handler(input string) (string, *int, error) {
	var err error
	components := strings.Split(input, " ")
	if len(components) > 1 {
		err = utils.ErrTooManyArguments
	}
	if vars.Prompt != "$ " {
		w.Output = strings.TrimSuffix(vars.Prompt, "$ ")
	} else {
		w.Output = "Anonymous"
	}
	return w.Output, nil, err
}
