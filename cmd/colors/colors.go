package colors

import "github.com/fatih/color"

var WorkingDirectoryColor = color.New(color.FgBlue).SprintFunc()

var ErrorColor = color.New(color.FgRed).SprintFunc()

var UserColor = color.New(color.FgHiGreen).SprintFunc()
