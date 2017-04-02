package parse

import (
	"log"
	"unicode"
	"strings"
)

func MapFromHeaderLines(lines []string) map[string]string {
	headerMap := make(map[string]string)

	// Name of the header (e.g. "Subject", "From")
	var key string

	for _, line := range lines {
		var value string

		if !unicode.IsSpace(rune(line[0])) {
			splitIndex := strings.Index(line, ":")

			if splitIndex != -1 {
				key = line[:splitIndex]
				value = strings.TrimSpace(line[splitIndex + 1:])
			} else {
				log.Println("Continuation line did not begin with whitespace as per https://tools.ietf.org/html/rfc2822")
				continue
			}
		} else {
			// Line began with whitespace; it contains content continued
			// from the previous line
			value = " " + strings.TrimSpace(line)
		}
		// Escape newline characters, could be integral to meaning
		headerMap[key] += strings.Replace(value, "\n", "\\n", -1)
	}

	return headerMap
}

