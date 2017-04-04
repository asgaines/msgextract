package output

import (
	"testing"
	"os"
	"bufio"
	"reflect"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

func TestWriteFieldsJSON(t *testing.T) {
	tmpDir, err := ioutil.TempDir("../test_files", "tmp")
	if err != nil {
		t.Error(err)
	}
	outputPath := filepath.Join(tmpDir, "output.json")
	fields := []string{"Subject", "From", "Date"}
	defer os.RemoveAll(tmpDir)

	cases := []struct {
		parsedHeaders []map[string]string
		jsonOutput []map[string]string
	}{
		{
			[]map[string]string{
				map[string]string{
					"Subject": "Urgent",
					"From": "ron@example.com",
					"Date": "01 Apr 2012, +0000",
				},
			},
			[]map[string]string{
				map[string]string{
					"Subject": "Urgent",
					"From": "ron@example.com",
					"Date": "01 Apr 2012, +0000",
				},
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Subject": "",
					"From": "",
					"Date": "",
				},
			},
			[]map[string]string{
				map[string]string{
					"Subject": "",
					"From": "",
					"Date": "",
				},
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Return-Path": "",
					"Received": "",
					"To": "",
				},
			},
			[]map[string]string{
				map[string]string{
					"Subject": "",
					"From": "",
					"Date": "",
				},
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Return-Path": "",
					"Received": "",
					"Subject": "From a Secret Admirer",
				},
			},
			[]map[string]string{
				map[string]string{
					"Subject": "From a Secret Admirer",
					"From": "",
					"Date": "",
				},
			},
		},
	}

	for _, c := range cases {
		WriteFields(outputPath, c.parsedHeaders, fields, "json")

		reader, err := ioutil.ReadFile(outputPath)
		if err != nil {
			t.Error(err)
		}

		var results []map[string]string

		json.Unmarshal(reader, &results)

		if !reflect.DeepEqual(results, c.jsonOutput) {
			t.Errorf("Received %v, wanted %v", results, c.jsonOutput)
		}
	}
}

func TestWriteFieldsTSV(t *testing.T) {
	tmpDir, err := ioutil.TempDir("../test_files", "tmp")
	if err != nil {
		t.Error(err)
	}
	outputPath := filepath.Join(tmpDir, "output.tsv")
	fields := []string{"Subject", "From", "Date"}
	defer os.RemoveAll(tmpDir)

	cases := []struct {
		parsedHeaders []map[string]string
		tsvOutputLines []string
	}{
		{
			[]map[string]string{
				map[string]string{
					"Subject": "Urgent",
					"From": "ron@example.com",
					"Date": "01 Apr 2012, +0000",
				},
			},
			[]string{
				"Subject\tFrom\tDate",
				"Urgent\tron@example.com\t01 Apr 2012, +0000",
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Subject": "",
					"From": "",
					"Date": "",
				},
			},
			[]string{
				"Subject\tFrom\tDate",
				"\t\t",
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Return-Path": "",
					"Received": "",
					"To": "",
				},
			},
			[]string{
				"Subject\tFrom\tDate",
				"\t\t",
			},
		},
		{
			[]map[string]string{
				map[string]string{
					"Return-Path": "",
					"Received": "",
					"Subject": "From a Secret Admirer",
				},
			},
			[]string{
				"Subject\tFrom\tDate",
				"From a Secret Admirer\t\t",
			},
		},
	}

	for _, c := range cases {
		WriteFields(outputPath, c.parsedHeaders, fields, "tsv")

		reader, err := os.Open(outputPath)
		if err != nil {
			t.Error(err)
		}

		scanner := bufio.NewScanner(reader)

		for _, line := range c.tsvOutputLines {
			scanner.Scan()
			resultLine := scanner.Text()

			if line != resultLine {
				t.Errorf("Received %v, wanted %v", resultLine, line)
			}
		}
	}
}

