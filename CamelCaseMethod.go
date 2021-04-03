package kata

import (
	"strings"
)

func CamelCase(s string) string {
	return strings.ReplaceAll(strings.Title(s), " ", "")
}
