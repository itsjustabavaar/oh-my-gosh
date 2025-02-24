package auth

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"golang.org/x/term"
	"strings"
	"syscall"
)

type LoginCommand struct {
	Input  string
	Output string
}

func (m *LoginCommand) Handler() (string, *int, error) {
	var err error
	name := strings.TrimSpace(strings.TrimPrefix(m.Input, "login "))
	if name == "" {
		err = utils.ErrInvalidUsername
	} else {
		vars.Prompt = name + "$ "
		m.Output = "login succeed"
	}
	return m.Output, nil, err
}

func PasswordReader() (string, error) {
	var password string
	fmt.Print("enter password: ")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		err = utils.ErrReadingPassword
	} else {
		password = string(bytePassword)
	}
	return password, err
}
