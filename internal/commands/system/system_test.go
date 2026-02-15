package system_test

import (
	"bytes"
	"os"
	"runtime"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
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

	handler.InputHandler(systemCommand)
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

	if output == "" {
		t.Fatal("output should not be empty")
	}
}
