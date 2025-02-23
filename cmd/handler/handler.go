package handler

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/auth"
	"github.com/itsjustabavaar/oh-my-gosh/internal/exit"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"os"
	"strings"
)

func HandleInput(input string, prompt *string) (state utils.HandlerState, output string, err error) {
	if strings.HasPrefix(input, "mynameis") {
		return auth.MyNameIsHandler(input, prompt)
	} else if input == "whoami" {
		return auth.WhoAmIHandler(prompt)
	} else if strings.HasPrefix(input, "exit") {
		state, exitCode, err := exit.Handler(input)
		if err != nil {
			return state, "", err
		}
		return
	}
	return utils.StateUnknown, input, nil
}
