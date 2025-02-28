package tests

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"runtime"
	"testing"
)

func TestTypeNotEnoughArguments(t *testing.T) {
	typeCommand := "type"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(typeCommand)
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

	output = utils.OutputCleaner(output)

	expectedOutput := "-gosh: type: not enough arguments"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestTypeBuiltins(t *testing.T) {
	typeCommand := "type cat"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(typeCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardOutput = oldStdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = utils.OutputCleaner(output)

	expectedOutput := "cat is a gosh builtin"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestTypeOsBuiltin(t *testing.T) {
	typeCommand := "type dir"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(typeCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardOutput = oldStdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = utils.OutputCleaner(output)

	var expectedOutput string = "-gosh: type: dir: command not found"

	if runtime.GOOS == "windows" {
		expectedOutput = "dir is your operating system builtin"
	}

	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestTypeUnknownCommand(t *testing.T) {
	typeCommand := "type fghueivooji"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(typeCommand)
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

	output = utils.OutputCleaner(output)

	expectedOutput := "-gosh: type: fghueivooji: command not found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestTypeSystemCommand(t *testing.T) {
	typeCommand := "type git"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(typeCommand)
	err := w.Close()
	if err != nil {
		return
	}
	vars.StandardOutput = oldStdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return
	}
	output := buf.String()

	output = utils.OutputCleaner(output)
	var expectedOutput string
	if runtime.GOOS == "windows" {
		expectedOutput = "git is C:\\Program Files\\Git\\cmd\\git.exe"
	} else {
		expectedOutput = "git is /usr/bin/git"
	}
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
