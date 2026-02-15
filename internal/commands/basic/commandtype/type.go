package commandtype

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

const typeCommand = "type"

type TypeCommand struct {
	Input  string
	Output string
}

func (t *TypeCommand) Execute() {
	components := util.SplitInput(t.Input, " ")
	if len(components) == 1 {
		util.PrintError(typeCommand, errorutil.ErrNotEnoughArguments)
		return
	}

	for i := 1; i < len(components); i++ {
		desiredType := components[i]
		if runtime.GOOS == "windows" {
			if _, ok := vars.WindowsBuiltins[desiredType]; ok {
				t.Output = fmt.Sprintf("%s is your operating system builtin", desiredType)
				util.PrintOutput(t.Output)
				return
			}
		}
		result, err := FindingType(desiredType)
		if err != nil {
			util.PrintError(typeCommand, err)
			continue
		}

		util.PrintOutput(result)
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
		return "", errorutil.ErrUnknownOsType
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

	return "", fmt.Errorf("%s: %w", inputType, errorutil.ErrCommandNotFound)
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
