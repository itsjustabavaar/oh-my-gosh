package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
)

const homeCommand = "homeCommand"

type HomeCommand struct {
	Output string
}

func (h *HomeCommand) Handler() {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		utils.Error(homeCommand, err)
		return
	}

	_, _ = fmt.Fprintf(vars.StandardOutput, "-gosh: %s: is a directory, it's your homeCommand :)\n", homeDirectory)
}
