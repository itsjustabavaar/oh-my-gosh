package basiccommands

import (
	"errors"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

const cdCommand = "cd"

type CdCommand struct {
	Input  string
	Output string
}

func (c *CdCommand) Execute() {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		utils.PrintError(cdCommand, err)
		return
	}

	workingDirectory := &vars.CurrentWorkingDirectory

	components := utils.SplitInput(c.Input, " ")
	if len(components) > 2 {
		utils.PrintError(cdCommand, utils.ErrTooManyArguments)
		return
	}

	if len(components) == 1 {
		err = os.Chdir(homeDirectory)
		if err != nil {
			utils.PrintError(cdCommand, err)
		}
		return
	}

	destinationDirectory := components[1]
	switch {

	case destinationDirectory == "-":
		previousDirectory := os.Getenv("OLDPWD")
		if previousDirectory == "" {
			previousDirectory = *workingDirectory
		}
		destinationDirectory = previousDirectory

	case strings.HasPrefix(destinationDirectory, "~"):
		destinationDirectory = strings.ReplaceAll(destinationDirectory, "~", homeDirectory)

	default:
		pathComponents := strings.Split(destinationDirectory, "/")
		var pathBuilder string
		for _, component := range pathComponents {
			pathBuilder = filepath.Join(pathBuilder, component)
			info, statErr := os.Stat(pathBuilder)
			if statErr == nil && !info.IsDir() {
				err = fmt.Errorf("%s is not a directory", strings.Join(pathComponents, "/"))
				utils.PrintError(cdCommand, err)
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
				utils.PrintError(cdCommand, utils.ErrDirectoryNotExists)
				return
			case errors.Is(pathErr.Err, syscall.EACCES):
				utils.PrintError(cdCommand, utils.ErrPermissionDenied)
				return
			default:
				utils.PrintError(cdCommand, utils.ErrChangingDirectory)
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
		utils.PrintError(cdCommand, err)
		return
	}

	*workingDirectory = destinationDirectory

	utils.PrintOutput(c.Output)
}
