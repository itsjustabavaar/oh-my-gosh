package whoami

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

type WhoAmICommand struct {
	Output string
}

func (w *WhoAmICommand) Execute() {
	w.Output = "Anonymous"

	if vars.CurrentUser.Username != "" {
		w.Output = vars.CurrentUser.Username
	}

	util.PrintOutput(w.Output)
}
