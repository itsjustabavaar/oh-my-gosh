package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"io"
	"os"
)

const catCommand = "cat"

type CatCommand struct {
	Input  string
	Output string
}

func (c *CatCommand) Execute() {
	components := utils.SplitInput(c.Input, " ")

	if len(components) == 1 {
		utils.PrintError(catCommand, utils.ErrNotEnoughArguments)
		return
	}

	for i := 1; i < len(components); i++ {
		fileContent, catErr := catFile(components[i])
		if catErr != nil {
			utils.PrintError(catCommand, fmt.Errorf("%s: %s", components[i], catErr))
			continue
		}
		utils.PrintOutput(fileContent)
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
			return fileOutput, utils.ErrFileNotFound
		}

		return fileOutput, utils.ErrFileOpening
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	data, _ := io.ReadAll(file)

	fileOutput = string(data)

	return fileOutput, nil
}
