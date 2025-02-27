package main

import (
	"bufio"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/colors"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/internal/database"
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	workingDirectory := &vars.CurrentWorkingDirectory

	utils.HandleInterrupt()

	err := models.MigrateDB(database.GetDB())
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
		os.Exit(1)
	}

	for {
		printPrompt(workingDirectory)

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "" {
			continue
		}

		err = basiccommands.StoreCommandHistory(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
			os.Exit(1)
		}

		if strings.HasPrefix(input, "clear") || strings.HasPrefix(input, "cls") {
			err := utils.ClearScreen()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
				os.Exit(1)
			}
			continue
		}

		handler.InputHandler(input)
	}
}

func printPrompt(workingDirectory *string) {
	*workingDirectory = utils.GetCurrentDirectory()
	fmt.Printf("%s:%s:%s ", colors.WorkingDirectoryColor(*workingDirectory), colors.UserColor(vars.CurrentUser.Username), vars.Prompt)
}
