package pwd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func TestPwd(t *testing.T) {
	currentDirectory, _ := os.Getwd()
	pwdCommand := "pwd"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(pwdCommand)
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

	if output != currentDirectory {
		t.Fatalf("Expected: %s, Actual: %s", currentDirectory, output)
	}
}
