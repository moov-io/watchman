package api

import (
	"strings"
)

var (
	newlineRemover = strings.NewReplacer("\n", "", "\r", "")
)

func CleanUserInput(input string) string {
	return newlineRemover.Replace(input)
}
