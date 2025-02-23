package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func MyNameIsHandler(input string, prompt *string) (utils.HandlerState, string, error) {
	name := strings.TrimSpace(strings.TrimPrefix(input, "mynameis "))
	if name != "" {
		*prompt = name + "$ "
		return utils.StateMyNameIs, "", nil
	}
	return utils.StateMyNameIs, "", utils.ErrInvalidUsername
}
