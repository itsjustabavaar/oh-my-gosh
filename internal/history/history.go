package history

type HistoryCommand struct {
	Output string
}

func (h *HistoryCommand) Handler() (string, *int, error) {
	return "", nil, nil
}
