package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"runtime"
	"testing"
)

func TestSystemCommands(t *testing.T) {
	var systemCommand string

	if runtime.GOOS == "windows" {
		systemCommand = "dir"
	} else {
		systemCommand = "ls"
	}

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(systemCommand)
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

	if output == "" {
		t.Fatal("output should not be empty")
	}
}
