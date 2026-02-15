package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/itsjustabavaar/oh-my-gosh/internal/colors"
	"github.com/itsjustabavaar/oh-my-gosh/internal/commands/basic/history"
	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util/dbutil"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	util.HandleInterrupt()

	db, err := dbutil.GormDB()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
		os.Exit(1)
	}

	err = models.MigrateDB(db)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
		os.Exit(1)
	}

	for {
		printPrompt()

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if input == "" {
			continue
		}

		err = history.StoreCommandHistory(input)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
			os.Exit(1)
		}

		if strings.HasPrefix(input, "clear") || strings.HasPrefix(input, "cls") {
			err := util.ClearScreen()
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
				os.Exit(1)
			}
			continue
		}

		handler.InputHandler(input)
	}
}

func printPrompt() {
	workingDirectory := util.GetCurrentDirectory()
	fmt.Printf("%s:%s:%s ", colors.WorkingDirectoryColor(workingDirectory), colors.UserColor(vars.CurrentUser.Username), vars.Prompt)
}
