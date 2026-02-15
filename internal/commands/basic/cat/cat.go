package cat

import (
	"fmt"
	"io"
	"os"

	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/errorutil"
)

const catCommand = "cat"

type CatCommand struct {
	Input  string
	Output string
}

func (c *CatCommand) Execute() {
	components := util.SplitInput(c.Input, " ")

	if len(components) == 1 {
		util.PrintError(catCommand, errorutil.ErrNotEnoughArguments)
		return
	}

	for i := 1; i < len(components); i++ {
		fileContent, catErr := catFile(components[i])
		if catErr != nil {
			util.PrintError(catCommand, fmt.Errorf("%s: %s", components[i], catErr))
			continue
		}
		util.PrintOutput(fileContent)
	}
}

func catFile(filename string) (string, error) {
	var fileOutput string

	info, statErr := os.Stat(filename)
	if statErr == nil && info.IsDir() {
		err := fmt.Errorf("is a directory")
		return "", err
	}

	file, openErr := os.Open(filename)

	if openErr != nil {
		if os.IsNotExist(openErr) {
			return fileOutput, errorutil.ErrFileNotFound
		}

		return fileOutput, errorutil.ErrFileOpening
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	data, _ := io.ReadAll(file)

	fileOutput = string(data)

	return fileOutput, nil
}
