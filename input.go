package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	trim := strings.TrimSpace(lower)
	fields := strings.Fields(trim)

	return fields
}
