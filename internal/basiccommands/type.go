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
	//if len(components) > 2 {
	//	utils.Error(typeCommand, utils.ErrTooManyArguments)
	//	return
	//}
	for i := 1; i < len(components); i++ {
		desiredType := components[i]
		result, err := FindingType(desiredType)
		if err != nil {
			utils.Error(typeCommand, err)
			continue
		}
		_, _ = fmt.Fprintln(os.Stdout, result)
	}
}

func FindingType(inputType string) (string, error) {
	desiredType := inputType
	if _, ok := vars.GoshBuiltins[desiredType]; ok {
		return fmt.Sprintf("%s is a gosh builtin", desiredType), nil
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
		return "", utils.ErrUnknownOsType
	}
	for _, path := range pathEnvComponents {
		ok, err := CheckFileExistence(desiredType, path)
		if err != nil {
			return "", err
		}
		if ok {
			return fmt.Sprintf("%s is %s", inputType, filepath.Join(path, desiredType)), nil
		}
	}
	return "", fmt.Errorf("%s: %w", inputType, utils.ErrCommandNotFound)
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
