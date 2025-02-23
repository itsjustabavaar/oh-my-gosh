package auth

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strings"
)

func WhoAmIHandler(prompt *string) (utils.HandlerState, string, error) {
	if *prompt != "$ " {
		return utils.StateWhoAmI, strings.TrimSuffix(*prompt, "$ "), nil
	}
	return utils.StateWhoAmI, "Anonymous", nil
}
