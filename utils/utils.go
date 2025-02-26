package utils

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"golang.org/x/term"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func SplitInput(input string, separationMethod string) []string {
	parsedComponents := make([]string, 0)
	components := strings.Split(input, separationMethod)
	for _, component := range components {
		if component != "" {
			parsedComponents = append(parsedComponents, component)
		}
	}
	return parsedComponents
}

func Error(commandName string, err error) {
	err = fmt.Errorf("-gosh: %s: %w", commandName, err)
	_, _ = fmt.Fprintln(vars.StandardError, err)
}

func PasswordReader(prompt string) (string, error) {
	var password string

	fmt.Printf("%s: ", prompt)

	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return password, ErrReadingPassword
	}

	password = string(bytePassword)

	fmt.Println()

	return password, err
}

func HandleInterrupt(originalState *term.State) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		_ = term.Restore(int(syscall.Stdin), originalState)
		os.Exit(1)
	}()
}

func GetCurrentDirectory() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		_ = fmt.Errorf("unable to determine current directory")
	}

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("unable to determine home directory")
	}

	if strings.HasPrefix(currentDirectory, userHomeDirectory) {
		currentDirectory = strings.ReplaceAll(currentDirectory, userHomeDirectory, "~")
	}

	return currentDirectory
}
