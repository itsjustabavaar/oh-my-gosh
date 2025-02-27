package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"io"
	"os"
)

const catCommand = "cat"

type CatCommand struct {
	Input  string
	Output string
}

func (c *CatCommand) Handler() {
	components := utils.SplitInput(c.Input, " ")
	if len(components) == 1 {
		utils.Error(catCommand, utils.ErrNotEnoughArguments)
		return
	}

	for i := 1; i < len(components); i++ {
		fileContent, catErr := catFile(components[i])
		if catErr != nil {
			utils.Error(catCommand, fmt.Errorf("%s: %s", components[i], catErr))
		} else {
			_, _ = fmt.Fprintln(vars.StandardOutput, fileContent)
		}
	}
}

func catFile(filename string) (string, error) {
	var fileOutput string

	file, openErr := os.Open(filename)
	if openErr != nil {
		if os.IsNotExist(openErr) {
			return fileOutput, utils.ErrFileNotFound
		} else {
			return fileOutput, utils.ErrFileOpening
		}
	}

	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {
			_ = fmt.Errorf(utils.ErrFileClosing.Error())
		}
	}(file)

	data, readErr := io.ReadAll(file)
	if readErr != nil {
		return fileOutput, utils.ErrFileReading
	}

	fileOutput = string(data)

	return fileOutput, nil
}
