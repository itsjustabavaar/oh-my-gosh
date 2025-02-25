package basiccommands

import (
	"fmt"
	"os"
)

type HomeCommand struct {
	Output string
}

func (h *HomeCommand) Handler() (string, *int, error) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", nil, err
	}

	h.Output = fmt.Sprintf("-gosh: %s: is a directory, it's your home :)", homeDirectory)

	return h.Output, nil, nil
}
