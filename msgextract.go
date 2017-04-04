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

	// Guard against invalid output file formats
	if !ValidFormats[outputFormat] {
		flag.Usage()
		os.Exit(1)
	}

	gzippedArchivePath := posArgs[0]

	tmpDir, err := ioutil.TempDir(".", "tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	archivePath := unpack.CreateArchiveName(tmpDir, gzippedArchivePath)
	outputPath := posArgs[1]

	err = unpack.Gzip(gzippedArchivePath, archivePath)
	if err != nil {
		log.Fatal(err)
	}

	// Channel to be fed the email header lines as they are
	// processed by tar function
	headerChan := make(chan []string)

	go func() {
		err = unpack.Tar(archivePath, headerChan)
		if err != nil {
			log.Fatal(err)
		}
		// Close the channel, releasing the blockage
		close(headerChan)
	}()

	// Parse through the header lines received through channel
	var parsedHeaders []map[string]string
	for headers := range headerChan {
		parsedHeaders = append(parsedHeaders, parse.MapFromHeaderLines(headers))
	}

	output.WriteFields(outputPath, parsedHeaders, fields, outputFormat)
}

