package tests

import (
	"bytes"
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestEchoSingleQuote(t *testing.T) {
	echoCommand := "echo 'salam'"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(echoCommand)
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
}

func TestEchoDoubleQuote(t *testing.T) {
	echoCommand := "echo \"$PATH ab \\a \\$ \\` \\\" \\\\\""

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(echoCommand)
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

	pathEnv := os.Getenv("PATH")

	expectedOutput := fmt.Sprintf("%s ab \\a $ ` \" \\", pathEnv)
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestEchoMultipleWithEnv(t *testing.T) {
	echoCommand := `echo salam chetori$PATH`

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(echoCommand)
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

	pathEnv := os.Getenv("PATH")

	expectedOutput := fmt.Sprintf("salam chetori%s", pathEnv)
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
