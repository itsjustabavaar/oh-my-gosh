package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestAnonymousWhoAmI(t *testing.T) {
	logoutCommand := "logout"

	InputHandler(logoutCommand)

	whoAmICommand := "whoami"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(whoAmICommand)
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

	expectedOutput := "Anonymous"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestUserWhoAmI(t *testing.T) {
	addUserCommand := "adduser testuser5 1234"

	InputHandler(addUserCommand)

	loginCommand := "login testuser5 1234"

	InputHandler(loginCommand)

	whoAmICommand := "whoami"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(whoAmICommand)
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

	expectedOutput := "testuser5"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
	err = user.DeleteUser("testuser5")
	if err != nil {
		t.Fatal("unexpected error")
	}
}
