package basiccommands

type TypeCommand struct {
	Input  string
	Output string
}

func (t *TypeCommand) Execute() (string, *int, error) {
	return "", nil, nil
}
