package wsgoutils

import (
	"fmt"
	"regexp"
	"strings"
)

// Match https://localhost:3000 or something.something.com or something.com
var remotehostPattern = regexp.MustCompile(`((?:(?:ht|f)tps?:\/\/)?(?:(?:(?:[^\/$\s]+\.)?[^\/$\s]+\.[^\/$\s]+)|(?:localhost:[\d]+)))`)

// Matches the optional protocol, host and optional port number
func MatchRemoteHost(fullPath string) []string {
	strMatches := fmt.Sprintf("%q\n", remotehostPattern.FindAll([]byte(fullPath), -1))
	// ["something.com"] => "[", "something.com", "]"
	matches := strings.Split(strMatches, "\"")
	actualMatches := []string{}
	for _, match := range matches {
		if match != "[" && match != "]" {
			actualMatches = append(actualMatches, match)
		}
	}
	return actualMatches
}
