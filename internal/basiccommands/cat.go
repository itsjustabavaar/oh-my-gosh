package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/colors"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"io"
	"os"
	"strings"
)

type CatCommand struct {
	Input  string
	Output string
}

func (c *CatCommand) Handler() (string, *int, error) {
	catResult := make([]string, 0)

	components := strings.Split(c.Input, " ")
	if len(components) == 1 {
		return c.Output, nil, utils.ErrNoFileProvided
	}

	for i := 1; i < len(components); i++ {
		fileContent, catErr := catFile(components[i])
		if catErr != nil {
			catResult = append(catResult, fmt.Sprintf("Error: cat: %s: %s", components[i], colors.ErrorColor(catErr)))
		} else {
			catResult = append(catResult, fmt.Sprintf("%s", fileContent))
		}
	}
	return strings.Join(catResult, vars.CatSep), nil, nil
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
