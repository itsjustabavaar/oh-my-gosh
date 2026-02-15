package vars

import (
	"io"
	"os"
	"time"

	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
)

var (
	StandardOutput io.Writer = os.Stdout
	StandardError  io.Writer = os.Stderr
)

var Prompt = "$"

type History struct {
	Command   string
	Timestamp time.Time
}

var AnonymousHistory = make([]History, 10)

var (
	CurrentUser             = &models.User{}
	CurrentWorkingDirectory string
)

var GoshBuiltins = map[string]struct{}{
	"gosh":    {},
	"cat":     {},
	"cd":      {},
	"login":   {},
	"history": {},
	"logout":  {},
	"whoami":  {},
	"adduser": {},
	"echo":    {},
	"exit":    {},
	"pwd":     {},
	"type":    {},
}

var WindowsBuiltins = map[string]struct{}{
	"dir":   {},
	"mkdir": {},
	"rmdir": {},
	"del":   {},
	"cls":   {},
	"copy":  {},
	"move":  {},
	"ren":   {},
	"ping":  {},
}

var RedirectionSigns = map[string]struct{}{
	">":   {},
	">>":  {},
	"1>":  {},
	"1>>": {},
	"2>":  {},
	"2>>": {},
	"&>":  {},
	"&>>": {},
}
