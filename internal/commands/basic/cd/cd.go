package cd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

const cdCommand = "cd"

type CdCommand struct {
	Input  string
	Output string
}

func (c *CdCommand) Execute() {
	homeDirectory, _ := os.UserHomeDir()

	workingDirectory := &vars.CurrentWorkingDirectory

	components := util.SplitInput(c.Input, " ")

	componentsCount := len(components)

	switch {

	case componentsCount == 1:
		_ = os.Chdir(homeDirectory)
		*workingDirectory = homeDirectory
		return

	case componentsCount > 2:
		util.PrintError(cdCommand, errorutil.ErrTooManyArguments)
		return

	default:

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
				err := fmt.Errorf("%s is not a directory", strings.Join(pathComponents, "/"))
				util.PrintError(cdCommand, err)
				return
			}
		}

	}

	err := os.Chdir(destinationDirectory)
	if err != nil {
		var pathErr *os.PathError

		if errors.As(err, &pathErr) {
			switch {
			case errors.Is(pathErr.Err, syscall.ENOENT):
				util.PrintError(cdCommand, errorutil.ErrDirectoryNotExists)
				return
			default:
				util.PrintError(cdCommand, errorutil.ErrChangingDirectory)
				return
			}
		}
	}

	previousDirectory := *workingDirectory

	_ = os.Setenv("OLDPWD", previousDirectory)

	currentDirectory, _ := os.Getwd()

	absPath, _ := filepath.Abs(currentDirectory)

	*workingDirectory = absPath

	util.PrintOutput(c.Output)
}
