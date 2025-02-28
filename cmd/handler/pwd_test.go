package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
	currentDirectory, _ := os.Getwd()
	pwdCommand := "pwd"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(pwdCommand)
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

	if output != currentDirectory {
		t.Fatalf("Expected: %s, Actual: %s", currentDirectory, output)
	}
}
