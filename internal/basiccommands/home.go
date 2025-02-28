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
	homeDirectory, _ := os.UserHomeDir()

	h.Output = fmt.Sprintf("-gosh: %s: is a directory, it's your homeCommand :)\n", homeDirectory)

	utils.PrintOutput(h.Output)
}
