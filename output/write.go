package output

import (
	"os"
	"log"
	"strings"
	"encoding/json"
)

func WriteFields(outputPath string,
		parsedHeaders []map[string]string,
		fields []string,
		format string) {
	writer, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	switch format {
	case "json":
		// Create new slice of maps to store required data
		var allFieldHeaders []map[string]string

		for _, headers := range parsedHeaders {
			fieldHeaders := make(map[string]string)

			for _, field := range fields {
				fieldHeaders[field] = headers[field]
			}
			allFieldHeaders = append(allFieldHeaders, fieldHeaders)
		}
		json.NewEncoder(writer).Encode(allFieldHeaders)
	case "tsv":
		// Write description of fields on first line (table header)
		writer.WriteString(strings.Join(fields, "\t") + "\n")

		for _, headers := range parsedHeaders {
			var content []string

			for _, field := range fields {
				content = append(content, headers[field])
			}
			writer.WriteString(strings.Join(content, "\t") + "\n")
		}
	}
}

