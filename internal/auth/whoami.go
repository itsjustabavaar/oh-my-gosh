package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type WhoAmICommand struct {
	Output string
}

func (w *WhoAmICommand) Handler() (string, *int, error) {
	if vars.CurrentUser.Username != "" {
		w.Output = vars.CurrentUser.Username
	} else {
		w.Output = "Anonymous"
	}

	return w.Output, nil, nil
}
