package auth

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type WhoAmICommand struct {
	Output string
}

func (w *WhoAmICommand) Handler() {
	if vars.CurrentUser.Username != "" {
		w.Output = vars.CurrentUser.Username
	} else {
		w.Output = "Anonymous"
	}

	_, _ = fmt.Fprintln(vars.StandardOutput, w.Output)
}
