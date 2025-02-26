package vars

import (
	"github.com/itsjustabavaar/oh-my-gosh/internal/models"
)

var Prompt = "$"

type History struct {
	Timestamp uint
	Count     uint
}

var AnonymousHistory = make(map[string]History)

var CurrentUser = &models.User{}
var CurrentWorkingDirectory string

var CatSep = "\n-----------------------------------------------------\n"
