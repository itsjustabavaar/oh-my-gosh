package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

type WhoAmIState struct {
	Output string
}

func (w *WhoAmIState) Handler(input string, prompt *string) (output string, err error) {
	components := strings.Split(input, " ")
	if len(components) > 1 {
		err = utils.ErrTooManyArguments
	}
	if *prompt != "$ " {
		w.Output = strings.TrimSuffix(*prompt, "$ ")
	} else {
		w.Output = "Anonymous"
	}
	return w.Output, err
}
