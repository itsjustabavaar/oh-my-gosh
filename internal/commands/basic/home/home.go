package home

import (
	"fmt"
	"os"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
)

const homeCommand = "homeCommand"

type HomeCommand struct {
	Output string
}

func (h *HomeCommand) Execute() {
	homeDirectory, _ := os.UserHomeDir()

	h.Output = fmt.Sprintf("-gosh: %s: is a directory, it's your homeCommand :)\n", homeDirectory)

	util.PrintOutput(h.Output)
}
