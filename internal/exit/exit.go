package exit

import (
	"github.com/itsjustabavaar/oh-my-gosh/utils"
	"strconv"
	"strings"
)

func Handler(input string) (state utils.HandlerState, output int, err error) {
	components := strings.Split(input, " ")
	if len(components) > 2 {
		return utils.StateExit, 0, utils.ErrTooManyArguments
	}
	output, err = strconv.Atoi(components[1])
	if err != nil {
		return utils.StateExit, 0, utils.ErrInvalidExitCode
	}
	return utils.StateExit, output, nil
}
