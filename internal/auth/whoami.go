package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

type WhoAmICommand struct {
	Output string
}

func (w *WhoAmICommand) Handler() {
	w.Output = "Anonymous"

	if vars.CurrentUser.Username != "" {
		w.Output = vars.CurrentUser.Username
	}

	utils.PrintOutput(w.Output)
}
