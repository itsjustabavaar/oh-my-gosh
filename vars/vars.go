package vars

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
	"time"
)

var Prompt = "$"

type History struct {
	Command   string
	Timestamp time.Time
}

var AnonymousHistory = make([]History, 10)

var CurrentUser = &models.User{}
var CurrentWorkingDirectory string

var CatSep = "\n-----------------------------------------------------\n"
