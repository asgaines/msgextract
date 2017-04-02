package unpack

import (
	"os"
	"io"
	"bufio"
	"strings"
	"compress/gzip"
	"archive/tar"
)

func Gzip(gzippedArchivePath, targetPath string) error {
	// Open the gzipped file for reading
	reader, err := os.Open(gzippedArchivePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Unpack the zipped reader into tar archive
	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	// Prepare file to store archive
	writer, err := os.Create(targetPath)
	if err != nil {
		return nil
	}
	defer writer.Close()

	// Write archive bytes to file
	_, err = io.Copy(writer, archive)
	return err
}

func Tar(tarPath string) (error, [][]string) {
	// Create structure to hold all maps of parsed email headers
	var allHeaderLines [][]string
	// Open tar file for reading
	reader, err := os.Open(tarPath)
	if err != nil {
		return err, allHeaderLines
	}
	defer reader.Close()


	tarReader := tar.NewReader(reader)

	// Iterate through all messages
	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err, allHeaderLines
		}

		// Only handle MSG files
		if tarHeader.Name[len(tarHeader.Name) - 3:] != "msg" {
			continue
		}

		scanner := bufio.NewScanner(tarReader)

		// Initialize new slice of strings to collect lines of the header
		var fileHeaderLines []string

		// Load the slice with header lines
		for scanner.Scan() {
			line := scanner.Text()
			// Formatting specified at https://tools.ietf.org/html/rfc2822
			if strings.TrimSpace(line) == "" {
				// End of header section
				break
			} else {
				fileHeaderLines = append(fileHeaderLines, line)
			}
		}

		allHeaderLines = append(allHeaderLines, fileHeaderLines)
	}

	return err, allHeaderLines
}

func CreateArchiveName(gzippedPath string) string {
	return gzippedPath[:len(gzippedPath) - 3]
}


