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

const cdCommand = "cd"

type CdCommand struct {
	Input  string
	Output string
}

func (c *CdCommand) Handler() {
	homeDirectory, err := os.UserHomeDir()
	workingDirectory := &vars.CurrentWorkingDirectory
	if err != nil {
		utils.Error(cdCommand, err)
		return
	}

	components := utils.SplitInput(c.Input, " ")
	if len(components) > 2 {
		utils.Error(cdCommand, utils.ErrTooManyArguments)
		return
	}

	if len(components) == 1 {
		err = os.Chdir(homeDirectory)
		if err != nil {
			utils.Error(cdCommand, err)
			return
		}
	}

	destinationDirectory := components[1]

	if destinationDirectory == "-" {
		previousDirectory := os.Getenv("OLDPWD")
		if previousDirectory == "" {
			previousDirectory = *workingDirectory
		}
		destinationDirectory = previousDirectory
	} else if strings.HasPrefix(destinationDirectory, "~") {
		destinationDirectory = strings.ReplaceAll(destinationDirectory, "~", homeDirectory)
	} else {
		pathComponents := strings.Split(destinationDirectory, "/")
		for _, component := range pathComponents {
			if strings.Contains(component, ".") {
				utils.Error(cdCommand, fmt.Errorf("%s is not a directory", strings.Join(pathComponents, "/")))
				return
			}
		}
	}

	err = os.Chdir(destinationDirectory)
	if err != nil {
		var pathErr *os.PathError

		if errors.As(err, &pathErr) {
			switch {
			case errors.Is(pathErr.Err, syscall.ENOENT):
				utils.Error(cdCommand, utils.ErrDirectoryNotExists)
				return
			case errors.Is(pathErr.Err, syscall.EACCES):
				utils.Error(cdCommand, utils.ErrPermissionDenied)
				return
			default:
				utils.Error(cdCommand, utils.ErrChangingDirectory)
				return
			}
		}
	}

	previousDirectory := *workingDirectory
	if strings.HasPrefix(previousDirectory, "~") {
		previousDirectory = strings.ReplaceAll(previousDirectory, "~", homeDirectory)
	}

	err = os.Setenv("OLDPWD", previousDirectory)
	if err != nil {
		utils.Error(cdCommand, err)
		return
	}

	*workingDirectory = destinationDirectory

	_, _ = fmt.Fprintln(vars.StandardOutput, c.Output)
}
