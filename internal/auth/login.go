package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"strings"
)

type MyNameIsState struct {
	Output string
}

func (m *MyNameIsState) Handler(input string) (string, *int, error) {
	var err error
	name := strings.TrimSpace(strings.TrimPrefix(input, "mynameis "))
	if name == "" {
		err = utils.ErrInvalidUsername
	} else {
		vars.Prompt = name + "$ "
		m.Output = "login succeed"
	}
	return m.Output, nil, err
}
