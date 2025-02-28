package tests

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/cmd/handler"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestLoginLogout(t *testing.T) {
	addUserCommand := "adduser testuser4 1234"

	handler.InputHandler(addUserCommand)

	loginCommand := "login testuser4 1234"

	handler.InputHandler(loginCommand)

	logoutCommand := "logout"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	handler.InputHandler(logoutCommand)
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

	expectedOutput := "logout succeed"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser4")
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func TestAlreadyLoggedOut(t *testing.T) {
	addUserCommand := "adduser testuser4 1234"

	handler.InputHandler(addUserCommand)

	loginCommand := "login testuser4 1234"

	handler.InputHandler(loginCommand)

	logoutCommand := "logout"

	handler.InputHandler(logoutCommand)

	oldStderr := vars.StandardError
	r, w, _ := os.Pipe()
	vars.StandardError = w

	handler.InputHandler(logoutCommand)
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

	expectedOutput := "-gosh: logout: already logged out"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser4")
	if err != nil {
		t.Fatal("unexpected error")
	}
}
