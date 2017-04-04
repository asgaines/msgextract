package parse

import (
	"log"
	"testing"
	"reflect"
	"io/ioutil"
)

func init() {
	// Deactivate the logging of requests to screen
	log.SetOutput(ioutil.Discard)
}

func TestParseHeaderLines(t *testing.T) {
	cases := []struct {
		lines []string
		headerMap map[string]string
	}{
		{
			[]string{"Subject: Calling all adventurers"},
			map[string]string{"Subject": "Calling all adventurers"},
		},
		{
			[]string{"Subject:We met yesterday"},
			map[string]string{"Subject": "We met yesterday"},
		},
		{
			[]string{""},
			map[string]string{},
		},
		{
			[]string{"Line with no header key"},
			map[string]string{},
		},
		{
			[]string{"Subject: \t \n "},
			map[string]string{"Subject": ""},
		},
		{
			[]string{"Subject: fullmoon/\noomlluf"},
			map[string]string{"Subject": "fullmoon/\\noomlluf"},
		},
		{
			[]string{"Subject:      ...Is this the right person?   "},
			map[string]string{"Subject": "...Is this the right person?"},
		},
		{
			[]string{"Subject:\tHello?\t"},
			map[string]string{"Subject": "Hello?"},
		},
		{
			[]string{"Subject: Hello: My name is Ron"},
			map[string]string{"Subject": "Hello: My name is Ron"},
		},
		{
			[]string{
				"Subject: Regarding our previous conversation, ",
				" I wanted to reach out to clarify something ",
				" (sorry this is such a long title)",
			},
			map[string]string{
				"Subject": "Regarding our previous conversation, I wanted to reach out to clarify something (sorry this is such a long title)",
			},
		},
		{
			[]string{
				"Subject: I know I can be long-winded, but thanks  \t",
				"\tfor dealing with my long email subjects",
			},
			map[string]string{
				"Subject": "I know I can be long-winded, but thanks for dealing with my long email subjects",
			},
		},
		{
			[]string{
				"Subject: For your eyes only",
				"Date: Fri, 1 Apr 2011 14:14:49 -0000",
			},
			map[string]string{
				"Subject": "For your eyes only",
				"Date": "Fri, 1 Apr 2011 14:14:49 -0000",
			},
		},
	}

	for _, c := range cases {
		if out := MapFromHeaderLines(c.lines); !reflect.DeepEqual(out, c.headerMap) {
			t.Errorf("%v returned %v, wanted %v", c.lines, out, c.headerMap)
		}
	}
}

