package handler

import (
	"bytes"
	"github.com/itsjustabavaar/oh-my-gosh/internal/basiccommands"
	"github.com/itsjustabavaar/oh-my-gosh/internal/user"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"testing"
)

func TestCleanHistory(t *testing.T) {
	historyCommand := "history"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(historyCommand)
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

	expectedOutput := "no commands found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestHistoryCleaning(t *testing.T) {
	historyCleaningCommand := "history clean"

	InputHandler(historyCleaningCommand)

	historyCommand := "history"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(historyCommand)
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

	expectedOutput := "no commands found"
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}
}

func TestAnonymousHistory(t *testing.T) {
	err := basiccommands.StoreCommandHistory("ls")
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	historyCommand := "history"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(historyCommand)
	err = w.Close()
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

	expectedOutput := "| ls | 1 | "
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}

}

func TestUserHistory(t *testing.T) {
	addUserCommand := "adduser testuser6"

	InputHandler(addUserCommand)

	loginCommand := "login testuser6"

	InputHandler(loginCommand)

	err := basiccommands.StoreCommandHistory("ls")
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	InputHandler("logout")

	InputHandler("login testuser6")

	historyCommand := "history"

	oldStdout := vars.StandardOutput
	r, w, _ := os.Pipe()
	vars.StandardOutput = w

	InputHandler(historyCommand)
	err = w.Close()
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

	expectedOutput := "| ls | 1 | "
	if output != expectedOutput {
		t.Fatalf("Expected: %s, Actual: %s", expectedOutput, output)
	}

	InputHandler("history clean")
	InputHandler("logout")
	err = user.DeleteUser("testuser6")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

}
