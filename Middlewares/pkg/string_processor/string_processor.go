package string_processor

import (
	"strings"
)

func SplitString(str string) []string {
	return strings.Split(str, " ")
}
