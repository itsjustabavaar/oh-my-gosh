package cd_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func TestCdEmpty(t *testing.T) {
	cdCommand := "cd"

	handler.InputHandler(cdCommand)

	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	if vars.CurrentWorkingDirectory != homeDirectory {
		t.Fatal("unexpected working directory")
	}
}

func TestCdTooManyArguments(t *testing.T) {
	catCommand := "cd arg1 arg2"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(catCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardError = oldStderr

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: cd: too many arguments"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestCdDash(t *testing.T) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}
	cdCommand := "cd -"

	handler.InputHandler(cdCommand)

	if vars.CurrentWorkingDirectory != currentDirectory {
		t.Fatal("unexpected working directory")
	}
}

func TestCdHome(t *testing.T) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	cdCommand := "cd ~"

	handler.InputHandler(cdCommand)

	if vars.CurrentWorkingDirectory != homeDirectory {
		t.Fatal("unexpected working directory")
	}
}

func TestCdWithTilda(t *testing.T) {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("Unexpected error", err)
	}

	cdCommand := "cd ~"

	handler.InputHandler(cdCommand)

	cdCommand = "cd Desktop"

	handler.InputHandler(cdCommand)

	currWor := vars.CurrentWorkingDirectory
	if currWor != filepath.Join(homeDirectory, "Desktop") {
		t.Fatal("unexpected working directory")
	}
}

func TestCdToAFile(t *testing.T) {
	cdCommand := "cd ~"

	handler.InputHandler(cdCommand)

	cdCommand = "cd Desktop"

	handler.InputHandler(cdCommand)

	catCommand := "cat salam > salam.txt"

	handler.InputHandler(catCommand)

	cdCommand = "cd salam.txt"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(cdCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardError = oldStderr

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = util.OutputCleaner(output)

	expectedOutput := fmt.Sprint("-gosh: cd: salam.txt is not a directory")

	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestCdDirectoryNotFound(t *testing.T) {
	cdCommand := "cd ~"

	handler.InputHandler(cdCommand)

	cdCommand = "cd Desktop"

	handler.InputHandler(cdCommand)

	cdCommand = "cd dueigvwdvnefivurvg"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(cdCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardError = oldStderr

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = util.OutputCleaner(output)

	expectedOutput := fmt.Sprint("-gosh: cd: directory not exists")

	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
