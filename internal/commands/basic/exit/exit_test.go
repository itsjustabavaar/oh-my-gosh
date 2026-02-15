package exit_test

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/internal/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/util"
	"github.com/itsjustabavaar/oh-my-gosh/internal/vars"
)

func TestExitTooManyArguments(t *testing.T) {
	exitCommand := "exit 1 2 3 4"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(exitCommand)
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

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: exit: too many arguments"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestFunctionExitsWithSpecificCode(t *testing.T) {
	expectedExitCode := 0

	testBinary, err := os.Executable()
	if err != nil {
		t.Fatalf("Could not get test binary path: %v", err)
	}

	cmd := exec.Command(testBinary, "-test.run=TestHelperProcessForExitCode")

	cmd.Env = append(os.Environ(),
		"GO_WANT_HELPER_PROCESS=1")

	err = cmd.Run()

	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		exitCode := exitError.ExitCode()
		if exitCode != expectedExitCode {
			t.Errorf("Expected exit code %d, got %d", expectedExitCode, exitCode)
		}
	}
}

func TestHelperProcessForExitCode(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	exitCommand := "exit"

	handler.InputHandler(exitCommand)

	t.Error("Expected function to exit but it didn't")
}

func TestFunctionExitsCodeZero(t *testing.T) {
	expectedExitCode := 0

	testBinary, err := os.Executable()
	if err != nil {
		t.Fatalf("Could not get test binary path: %v", err)
	}

	cmd := exec.Command(testBinary, "-test.run=TestHelperProcessForExitCodeZero")

	cmd.Env = append(os.Environ(),
		"GO_WANT_HELPER_PROCESS=1")

	err = cmd.Run()

	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		exitCode := exitError.ExitCode()
		if exitCode != expectedExitCode {
			t.Errorf("Expected exit code %d, got %d", expectedExitCode, exitCode)
		}
	}
}

func TestHelperProcessForExitCodeZero(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	exitCommand := "exit"

	handler.InputHandler(exitCommand)

	t.Error("Expected function to exit but it didn't")
}

func TestInvalidExitCode(t *testing.T) {
	exitCommand := "exit salam"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(exitCommand)
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

	output = util.OutputCleaner(output)

	expectedOutput := "-gosh: exit: invalid exit code"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}
