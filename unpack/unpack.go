package unpack

import (
	"os"
	"io"
	"bufio"
	"strings"
	"compress/gzip"
	"archive/tar"
	"path/filepath"
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

func Tar(tarPath string, headerChan chan []string) error {
	// Open tar file for reading
	reader, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer reader.Close()


	tarReader := tar.NewReader(reader)

	// Iterate through all messages
	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		// Only handle MSG files
		if tarHeader.Name[len(tarHeader.Name) - 3:] != "msg" {
			continue
		}

		scanner := bufio.NewScanner(tarReader)

		// Initialize new slice of strings to collect lines of the header
		var headerLines []string

		// Load the slice with header lines
		for scanner.Scan() {
			line := scanner.Text()
			// Formatting specified at https://tools.ietf.org/html/rfc2822
			if strings.TrimSpace(line) == "" {
				// End of header section
				break
			} else {
				headerLines = append(headerLines, line)
			}
		}

		// Feed lines through channel
		headerChan <- headerLines
	}

	return err
}

func CreateArchiveName(tmpDir, gzippedPath string) string {
	fileName := filepath.Base(gzippedPath)
	return filepath.Join(tmpDir, fileName[:len(fileName) - 3])
}

