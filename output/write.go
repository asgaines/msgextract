package output

import (
	"os"
	"log"
	"strings"
)

func WriteFields(outputPath string,
		parsedHeaders []map[string]string,
		fields []string) {
	writer, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	for _, headers := range parsedHeaders {
		var content []string

		for _, field := range fields {
			if value, ok := headers[field]; ok {
				content = append(content, value)
			}
		}
		writer.WriteString(strings.Join(content, "\t") + "\n")
	}
}

