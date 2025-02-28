package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestOverrideRedirections(t *testing.T) {
	fillFileCommand := "echo salam > salam.txt"
	InputHandler(fillFileCommand)

	fillFileCommand = "echo salam 1> salam.txt"
	InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(catCommand)
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

func TestAppendRedirection(t *testing.T) {
	fillFileCommand := "echo salam >> salam.txt"
	InputHandler(fillFileCommand)

	fillFileCommand = "echo salam 1>> salam.txt"
	InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(catCommand)
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
	InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(catCommand)
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
	InputHandler(fillFileCommand)

	fillFileCommand = "type hfwieuh 2>> salam.txt"
	InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(catCommand)
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
	InputHandler(fillFileCommand)

	fillFileCommand = "echo salam &>> salam.txt"
	InputHandler(fillFileCommand)

	catCommand := "cat salam.txt"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(catCommand)
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

	expectedOutput := "-gosh: type: hfwieuh: command not foundsalam"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = os.Remove("salam.txt")
	if err != nil {
		t.Fatal("unexpected error")
	}
}
