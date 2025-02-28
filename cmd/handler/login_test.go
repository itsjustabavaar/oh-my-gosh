package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
	"testing"

	"github.com/itsjustabavaar/oh-my-gosh/vars"
)

/*
1. statement coverage:
	- go test ./... -coverprofile='coverage.out'
	- go tool cover -html='coverage.out'
*/

func TestLoginEmptyPassword(t *testing.T) {
	addUserCommand := "adduser testuser3"

	InputHandler(addUserCommand)

	loginCommand := "login testuser3"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(loginCommand)
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

	expectedOutput := "login succeed"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser3")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestLoginEmptyUsername(t *testing.T) {
	loginCommand := "login"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	InputHandler(loginCommand)
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

	expectedOutput := "-gosh: login: please tell me who you are"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestLoginTooManyArguments(t *testing.T) {
	loginCommand := "login testuser3 1234 4321"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	InputHandler(loginCommand)
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

	expectedOutput := "-gosh: login: too many arguments"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestLoginWithPassword(t *testing.T) {
	addUserCommand := "adduser testuser3 1234"

	InputHandler(addUserCommand)

	loginCommand := "login testuser3 1234"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(loginCommand)
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

	expectedOutput := "login succeed"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser3")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestLoginUserNotFound(t *testing.T) {
	loginCommand := "login testuser3 1234"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	InputHandler(loginCommand)
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

	expectedOutput := "-gosh: login: user not found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestLoginInvalidPassword(t *testing.T) {
	addUserCommand := "adduser testuser3 1234"

	InputHandler(addUserCommand)

	loginCommand := "login testuser3 1235"

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	InputHandler(loginCommand)
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

	expectedOutput := "-gosh: login: invalid password"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser3")
	if err != nil {
		t.Fatal("unexpected error")
	}
}
