package redirection_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func TestOverrideRedirections(t *testing.T) {
	fillFileCommand := "echo salam > salam.txt"
	handler.InputHandler(fillFileCommand)

	fillFileCommand = "echo salam 1> salam.txt"
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

	output = util.OutputCleaner(output)

	expectedOutput := "salam"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestAppendRedirection(t *testing.T) {
	fillFileCommand := "echo salam >> salam.txt"
	handler.InputHandler(fillFileCommand)

	fillFileCommand = "echo salam 1>> salam.txt"
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

	output = util.OutputCleaner(output)

	expectedOutput := "salamsalam"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestOverrideErrorRedirection(t *testing.T) {
	fillFileCommand := "type hfwieuh 2> salam.txt"
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

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: type: hfwieuh: command not found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestAppendErrorRedirection(t *testing.T) {
	fillFileCommand := "type hfwieuh 2>> salam.txt"
	handler.InputHandler(fillFileCommand)

	fillFileCommand = "type hfwieuh 2>> salam.txt"
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

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: type: hfwieuh: command not found-gosh: type: hfwieuh: command not found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestBothRedirection(t *testing.T) {
	fillFileCommand := "type hfwieuh &> salam.txt"
	handler.InputHandler(fillFileCommand)

	fillFileCommand = "echo salam &>> salam.txt"
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

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: type: hfwieuh: command not foundsalam"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}
