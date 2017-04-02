package main

import (
	"os"
	"log"
	"github.com/asgaines/msgextract/unpack"
	"github.com/asgaines/msgextract/parse"
	"github.com/asgaines/msgextract/output"
)

func main() {
	fields := []string{"Date", "From", "Subject"}

	if len(os.Args) < 3 {
		log.Fatal("Supply path to Gzipped archive and to output file")
	}

	gzippedArchivePath := os.Args[1]
	archivePath := unpack.CreateArchiveName(gzippedArchivePath)
	outputPath := os.Args[2]

	err := unpack.Gzip(gzippedArchivePath, archivePath)
	if err != nil {
		log.Fatal(err)
	}

	err, headerLines := unpack.Tar(archivePath)
	if err != nil {
		log.Fatal(err)
	}

	var parsedHeaders []map[string]string
	for _, lines := range headerLines {
		parsedHeaders = append(parsedHeaders, parse.MapFromHeaderLines(lines))
	}

	output.WriteFields(outputPath, parsedHeaders, fields)
}
