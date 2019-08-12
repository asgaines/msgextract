# Email Message Header Extractor

Parse through gzipped tar archive and output the desired header information to file. Email header fields collected: `Subject`, `From`, `Date`.

## Program Walk-Through

`msgextract` first unpacks the gzip-compressed file into a tar archive in a temporary directory. The MSG files in the archive are iterated through, reading the header (ignoring the potentially large body), and passing the header lines through a channel to a consumer. The consumer parses the lines into a map (key: header field name, value: header field content). This holistic map allows for arbitrary field selection, which are currently set to `Subject`, `From`, and `Date`. The selected fields are filtered from the map and output to file, in `json` or `tsv` format. See `Suggested Improvements` below for feature ideas and bugs.

## Installation

- Install [Go](https://golang.org/doc/install)
- `go get -u github.com/asgaines/msgextract/...`

## Usage

- `msgextract [--format=(json|tsv)] gzipped-archive.tar.gz output.(json|tsv)`
- If `-format` not specified, default is `json` output (note: optional args must precede positional args)

### Examples

- `msgextract gzipped-archive.tar.gz output.json`
- `msgextract --format=tsv gzipped-archive.tar.gz output.tsv`

## Suggested Improvements

- Concurrency in mapping of header lines returned by `unpack.Tar`
- Standardize formatting of `Date` header information
- Remove data duplication in `output.WriteFields`
 - Move header field filtering to new procedure
- Specify which header fields to collect as optional args
- Allow for multiple values for same header fields (e.g. multiple `Received`s)
- Add `csv` format capability
 - Will require addressing `,` in header field values

## Testing

- `cd path/to/msgextract`
- `go test ./...`

