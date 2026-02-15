package logout

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

const logoutCommand = "logout"

type LogOutCommand struct {
	Output string
}

func (l *LogOutCommand) Execute() {
	if vars.CurrentUser.Username == "" {
		util.PrintError(logoutCommand, errorutil.ErrAlreadyLoggedOut)
		return
	}

	vars.CurrentUser = &models.User{}

	l.Output = "logout succeed"

	util.PrintOutput(l.Output)
}
