package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

const logoutCommand = "logout"

type LogOutCommand struct {
	Output string
}

func (l *LogOutCommand) Handler() {
	if vars.CurrentUser.Username == "" {
		utils.PrintError(logoutCommand, utils.ErrAlreadyLoggedOut)
		return
	}

	vars.CurrentUser = &models.User{}

	l.Output = "logout succeed"

	utils.PrintOutput(l.Output)
}
