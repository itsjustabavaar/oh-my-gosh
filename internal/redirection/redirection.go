package redirection

import (
	"fmt"
	"os"
	"strings"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func DetectAndApplyRedirection(inputCommand *string) *os.File {
	inputComponents := util.SplitInput(*inputCommand, " ")

	if len(inputComponents) > 2 {
		possibleRedirection := inputComponents[len(inputComponents)-2]
		possibleFilePath := inputComponents[len(inputComponents)-1]

		if _, ok := vars.RedirectionSigns[possibleRedirection]; ok {
			file, err := applyRedirection(possibleRedirection, possibleFilePath)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)

				return nil
			}

			*inputCommand = strings.Join(inputComponents[:len(inputComponents)-2], " ")
			return file

		}

	}

	return nil
}

func applyRedirection(inputRedirection string, inputFilePath string) (*os.File, error) {
	var file *os.File
	var err error

	switch inputRedirection {

	case ">", "1>":
		file, err = openFileRedirection(inputFilePath, "override")
		if err != nil {
			return nil, err
		}

		vars.StandardOutput = file

	case ">>", "1>>":
		file, err = openFileRedirection(inputFilePath, "append")
		if err != nil {
			return nil, err
		}

		vars.StandardOutput = file

	case "2>":
		file, err = openFileRedirection(inputFilePath, "override")
		if err != nil {
			return nil, err
		}

		vars.StandardError = file

	case "2>>":
		file, err = openFileRedirection(inputFilePath, "append")
		if err != nil {
			return nil, err
		}

		vars.StandardError = file

	case "&>":
		file, err = openFileRedirection(inputFilePath, "override")
		if err != nil {
			return nil, err
		}

		vars.StandardOutput, vars.StandardError = file, file

	case "&>>":
		file, err = openFileRedirection(inputFilePath, "append")
		if err != nil {
			return nil, err
		}

		vars.StandardOutput, vars.StandardError = file, file
	}

	return file, nil
}

func openFileRedirection(filepath string, openMode string) (*os.File, error) {
	if openMode == "override" {
		file, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
		return file, nil
	}

	if openMode == "append" {
		file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return nil, errorutil.ErrFileOpening
		}

		return file, nil
	}

	return nil, errorutil.ErrUnknownMode
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		return
	}
}

func RestoreStates(file *os.File) {
	defer CloseFile(file)

	if vars.StandardOutput != os.Stdout {
		vars.StandardOutput = os.Stdout
	}

	if vars.StandardError != os.Stderr {
		vars.StandardError = os.Stderr
	}
}
