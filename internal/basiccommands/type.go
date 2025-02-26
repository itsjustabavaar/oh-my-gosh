package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const typeCommand = "type"

type TypeCommand struct {
	Input  string
	Output string
}

func (t *TypeCommand) Handler() {
	components := utils.SplitInput(t.Input, " ")
	if len(components) == 1 {
		utils.Error(typeCommand, utils.ErrNotEnoughArguments)
		return
	}
	if len(components) > 2 {
		utils.Error(typeCommand, utils.ErrTooManyArguments)
		return
	}
	desiredType := components[1]
	if _, ok := vars.GoshBuiltins[desiredType]; ok {
		_, _ = fmt.Fprintf(vars.StandardOutput, "%s is a gosh builtin\n", desiredType)
		return
	}

	pathEnv := os.Getenv("PATH")
	pathEnvComponents := make([]string, 0)

	switch runtime.GOOS {
	case "windows":
		pathEnvComponents = strings.Split(pathEnv, ";")
		desiredType += ".exe"
	case "linux":
		pathEnvComponents = strings.Split(pathEnv, ":")
	default:
		utils.Error(typeCommand, utils.ErrUnknownOsType)
		return
	}
	for _, path := range pathEnvComponents {
		ok, err := CheckFileExistence(desiredType, path)
		if err != nil {
			utils.Error(typeCommand, err)
			return
		}
		if ok {
			_, _ = fmt.Fprintf(vars.StandardOutput, "%s is %s\n", components[1], filepath.Join(path, desiredType))
			return
		}
	}
	utils.Error(typeCommand, fmt.Errorf("%s: %w", components[1], utils.ErrCommandNotFound))
	return
}

func CheckFileExistence(filename, directory string) (bool, error) {
	filePath := filepath.Join(directory, filename)
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
