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
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	workingDirectory := &vars.CurrentWorkingDirectory

	oldState, err := term.GetState(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Failed to get terminal state:", err)
		return
	}

	utils.HandleInterrupt(oldState)

	err = models.MigrateDB(database.GetDB())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		*workingDirectory = utils.GetCurrentDirectory()
		fmt.Printf("%s:%s:%s ", colors.WorkingDirectoryColor(*workingDirectory), colors.UserColor(vars.CurrentUser.Username), vars.Prompt)

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		if !strings.HasPrefix(input, "history") {
			err = basiccommands.StoreCommandHistory(input)
		}
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "-gosh: ", err)
			os.Exit(1)
		}

		handler.InputHandler(input)

	}
}
