package tests

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestCat(t *testing.T) {
	fillFileCommand := "echo salam > salam.txt"

	handler.InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(catCommand)
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

	expectedOutput := "salam"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestCatNotEnoughArguments(t *testing.T) {
	catCommand := "cat"

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

	output = utils.OutputCleaner(output)

	expectedOutput := "-gosh: cat: not enough arguments"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestCatFileNotFound(t *testing.T) {
	catCommand := "cat fvwiev.efguy"

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

	output = utils.OutputCleaner(output)

	expectedOutput := "-gosh: cat: fvwiev.efguy: no such file or directory"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
