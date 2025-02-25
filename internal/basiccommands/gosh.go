package basiccommands

type GoshCommand struct {
	Output string
}

func (g *GoshCommand) Handler() (string, *int, error) {
	g.Output = "I'm Here! :)"
	return g.Output, nil, nil
}
