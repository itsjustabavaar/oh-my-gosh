package vars

var Prompt string = "$ "

type History struct {
	Timestamp uint
	Count     uint
}

var AnonymousHistory map[string]History = make(map[string]History)
