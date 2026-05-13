package rainbow

import (
	"regexp"
	"strings"
)

var ansiRE = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRE.ReplaceAllString(s, "")
}

func wrapANSI(s, code string) string {
	seq := "\x1b[" + code + "m"
	if !strings.Contains(s, "\x1b") {
		return seq + s + "\x1b[0m"
	}

	prefixEnd := 0
	for {
		loc := ansiRE.FindStringIndex(s[prefixEnd:])
		if len(loc) != 2 || loc[0] != 0 {
			break
		}
		prefixEnd += loc[1]
	}
	if prefixEnd > 0 {
		s = s[:prefixEnd] + seq + s[prefixEnd:]
	} else {
		s = seq + s
	}
	if !strings.HasSuffix(s, "\x1b[0m") {
		s += "\x1b[0m"
	}
	return s
}
