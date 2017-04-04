package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"io/ioutil"
	"github.com/asgaines/msgextract/unpack"
	"github.com/asgaines/msgextract/parse"
	"github.com/asgaines/msgextract/output"
)

func main() {
	fields := []string{"Date", "From", "Subject"}

	var ValidFormats = map[string]bool {
		"json": true,
		"tsv": true,
	}

	var outputFormat string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [opt args] gzipped-archive.tar.gz output.json\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&outputFormat, "format", "json", "Formatting for the output file. Valid options: json, tsv")

	flag.Parse()

	posArgs := flag.Args()

	if len(posArgs) < 2 {
		log.Fatal("Supply path to Gzipped archive and to output file")
		flag.Usage()
		os.Exit(1)
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

	output.WriteFields(outputPath, parsedHeaders, fields, outputFormat)
}

