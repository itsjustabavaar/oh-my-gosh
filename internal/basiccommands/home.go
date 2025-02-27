package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
)

const homeCommand = "homeCommand"

type HomeCommand struct {
	Output string
}

func (h *HomeCommand) Execute() {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		utils.PrintError(homeCommand, err)
		return
	}

	h.Output = fmt.Sprintf("-gosh: %s: is a directory, it's your homeCommand :)\n", homeDirectory)

	utils.PrintOutput(h.Output)
}
