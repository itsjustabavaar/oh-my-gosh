package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

type MyNameIsState struct {
	Output string
}

func (m *MyNameIsState) Handler(input string, prompt *string) (output string, err error) {
	name := strings.TrimSpace(strings.TrimPrefix(input, "mynameis "))
	if name == "" {
		err = utils.ErrInvalidUsername
	} else {
		*prompt = name + "$ "
		m.Output = "login succeed"
	}
	return m.Output, err
}
