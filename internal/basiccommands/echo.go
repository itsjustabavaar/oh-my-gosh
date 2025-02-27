package basiccommands

import (
	"fmt"
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"github.com/itsjustabavaar/oh-my-gosh/vars"
	"os"
	"strings"
)

type EchoCommand struct {
	Input  string
	Output string
}

func (e *EchoCommand) Handler() {
	echoArgument := strings.TrimPrefix(e.Input, "echo ")
	if strings.HasPrefix(echoArgument, "'") && strings.HasSuffix(echoArgument, "'") {
		_, _ = fmt.Fprintln(vars.StandardOutput, strings.ReplaceAll(echoArgument, "'", ""))
	}
	if strings.HasPrefix(echoArgument, "\"") && strings.HasSuffix(echoArgument, "\"") {
		_, _ = fmt.Fprintln(vars.StandardOutput, processDoubleQuotedStrings(echoArgument))
	}

	echoPhrases := utils.SplitInput(echoArgument, " ")

	for idx, phrase := range echoPhrases {
		if strings.Contains(phrase, "$") {

			echoComponents := strings.Split(phrase,
				"$")
			for i := 1; i < len(echoComponents); i++ {
				echoComponents[i] = os.Getenv(echoComponents[i])
			}

			phrase = strings.Join(echoComponents, "")
		}

		echoPhrases[idx] = phrase
	}

	_, _ = fmt.Fprintln(vars.StandardOutput, strings.Join(echoPhrases, " "))
}

func processDoubleQuotedStrings(input string) string {
	innerContent := input[1 : len(input)-1]

	result, escaped := "", false

	for i := 0; i < len(innerContent); i++ {
		c := innerContent[i]

		if escaped {
			if c == '$' || c == '`' || c == '"' || c == '\\' || c == '\n' {
				result += string(c)
			} else {
				result += "\\" + string(c)
			}
			escaped = false
		} else if c == '\\' {
			escaped = true
		} else if c == '$' {
			if i+1 < len(innerContent) {
				varStart := i + 1
				varEnd := varStart

				for varEnd < len(innerContent) &&
					(innerContent[varEnd] >= 'a' && innerContent[varEnd] <= 'z' ||
						innerContent[varEnd] >= 'A' && innerContent[varEnd] <= 'Z' ||
						innerContent[varEnd] >= '0' && innerContent[varEnd] <= '9' ||
						innerContent[varEnd] == '_') {
					varEnd++
				}

				if varEnd > varStart {
					varName := innerContent[varStart:varEnd]
					result += os.Getenv(varName)
					i = varEnd - 1
				} else {
					result += string(c)
				}
			} else {
				result += string(c)
			}
		} else {
			result += string(c)
		}
	}

	if escaped {
		result += "\\"
	}

	return result
}
