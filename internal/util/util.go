package util

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
	"golang.org/x/term"
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

func ClearScreen() error {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func PrintError(commandName string, err error) {
	err = fmt.Errorf("-gosh: %s: %w", commandName, err)
	_, _ = fmt.Fprintln(vars.StandardError, err)
}

func OutputCleaner(input string) string {
	return strings.ReplaceAll(input, "\n", "")
}

func PrintOutput(output string) {
	_, _ = fmt.Fprintln(vars.StandardOutput, output)
}

func HandleInterrupt() {
	originalState, err := term.GetState(int(syscall.Stdin))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "-gosh: failed to get terminal state: ", err)
		os.Exit(1)
	}
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		_ = term.Restore(int(syscall.Stdin), originalState)
		os.Exit(1)
	}()
}

func GetCurrentDirectory() string {
	var currentDirectory string

	if vars.CurrentWorkingDirectory == "" {
		vars.CurrentWorkingDirectory, _ = os.Getwd()
	}

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		_ = fmt.Errorf("unable to determine home directory")
	}

	if strings.HasPrefix(vars.CurrentWorkingDirectory, userHomeDirectory) {
		currentDirectory = strings.ReplaceAll(vars.CurrentWorkingDirectory, userHomeDirectory, "~")
	}

	return currentDirectory
}
