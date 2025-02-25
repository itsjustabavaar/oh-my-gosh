package basiccommands

import (
	"errors"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"strings"
	"syscall"
)

type CdCommand struct {
	Input  string
	Output string
}

func (c *CdCommand) Handler() (string, *int, error) {
	homeDirectory, err := os.UserHomeDir()
	workingDirectory := &vars.CurrentWorkingDirectory
	if err != nil {
		return "", nil, err
	}

	components := utils.SplitInput(c.Input, " ")
	if len(components) > 2 {
		return "", nil, utils.ErrTooManyArguments
	}

	if len(components) == 1 {
		err = os.Chdir(homeDirectory)
		if err != nil {
			return "", nil, err
		}
	}

	destinationDirectory := components[1]
	if strings.HasPrefix(destinationDirectory, "~") {
		destinationDirectory = strings.ReplaceAll(destinationDirectory, "~", homeDirectory)
	}

	pathComponents := strings.Split(destinationDirectory, "/")
	for _, component := range pathComponents {
		if strings.Contains(component, ".") {
			return "", nil, fmt.Errorf("%s is not a directory", strings.Join(pathComponents, "/"))
		}
	}

	err = os.Chdir(destinationDirectory)
	if err != nil {
		var pathErr *os.PathError

		if errors.As(err, &pathErr) {
			switch {
			case errors.Is(pathErr.Err, syscall.ENOENT):
				return "", nil, utils.ErrDirectoryNotExists
			case errors.Is(pathErr.Err, syscall.EACCES):
				return "", nil, utils.ErrPermissionDenied
			default:
				return "", nil, utils.ErrChangingDirectory
			}
		}
	}

	*workingDirectory = destinationDirectory

	return "", nil, nil
}
