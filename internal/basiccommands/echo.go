package basiccommands

import (
	"os"
	"strings"
)

type EchoCommand struct {
	Input  string
	Output string
}

func (e *EchoCommand) Handler() (string, *int, error) {
	var err error
	echoArgument := strings.TrimPrefix(e.Input, "echo ")
	if strings.HasPrefix(echoArgument, "'") && strings.HasSuffix(echoArgument, "'") {
		return strings.Trim(echoArgument, "'"), nil, err
	} else {
		echoPhrases := strings.Split(echoArgument, " ")
		for idx, phrase := range echoPhrases {
			if strings.Contains(phrase, "$") {
				echoComponents := strings.Split(phrase, "$")
				for i := 1; i < len(echoComponents); i++ {
					echoComponents[i] = os.Getenv(echoComponents[i])
				}
				phrase = strings.Join(echoComponents, "")
			}
			echoPhrases[idx] = phrase
		}
		return strings.Join(echoPhrases, " "), nil, err
	}
}
