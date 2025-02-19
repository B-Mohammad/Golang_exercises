package cmd

import (
	"strings"
)

func Say(name []string) string {
	if len(name) == 0 {
		return "Hello, World!"
	}

	return "Hello, " + strings.Join(name, " ") + "!"
}
