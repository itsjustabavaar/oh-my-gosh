package tests

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestGosh(t *testing.T) {
	goshCommand := "gosh"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(goshCommand)
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

	expectedOutput := "I'm Here! :)"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
