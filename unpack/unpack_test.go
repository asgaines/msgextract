package unpack

import (
	"testing"
	"os"
	"reflect"
	"path/filepath"
)

func TestUnGzippedArchiveDoesNotYetExist(t *testing.T) {
	unGzipDir := "../test_files/targzs"

	// Target path where contents of file to be stored 
	// after unpacking
	unzippedPath := filepath.Join(unGzipDir, "testEmails.tar")

	// Ensure unpacked tar file does not yet exist
	if _, err := os.Open(unzippedPath); !os.IsNotExist(err) {
		// The file exists
		t.Errorf("Unzipped file already exists: %v", unzippedPath)
	}
}

func TestUnGzipCreatesArchive(t *testing.T) {
	unGzipDir := "../test_files/targzs"

	// Target path where contents of file to be stored 
	// after unpacking
	unzippedPath := filepath.Join(unGzipDir, "testEmails.tar")
	// Gzipped filename will be same as unzipped archive with
	// .gz extension added
	zippedPath := unzippedPath + ".gz"

	Gzip(zippedPath, unzippedPath)

	defer os.Remove(unzippedPath)

	// Ensure unpacked tar file now exists
	if _, err := os.Open(unzippedPath); os.IsNotExist(err) {
		// The file doesn't exist
		t.Errorf("Unzipped file not successfully created: %v", unzippedPath)
	}
}

func TestTar(t *testing.T) {
	tarDir := "../test_files/tars"

	cases := []struct {
		tarPath string
		parsedHeaders [][]string
	}{
		{
			filepath.Join(tarDir, "subject_date_from.tar"),
			[][]string{
				{
					"From: \"Darty\" <infos@contact-darty.com>",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
				},
			},
		},
		{
			filepath.Join(tarDir, "dirdepth1.tar"),
			[][]string{
				{
					"From: \"Darty\" <infos@contact-darty.com>",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
				},
			},
		},
		{
			filepath.Join(tarDir, "return_x-orig_received.tar"),
			[][]string{
				{
					"Return-Path: <out-582911-B2C71BD37AF148CE9D728B61264F854D@mail.beliefnet.com>",
					"X-Original-To: beliefnet@cp.monitor1.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 447A219825C",
					"\tfor <beliefnet@cp.monitor1.returnpath.net>; Fri,  1 Apr 2011 10:32:42 -0600 (MDT)",
				},
			},
		},
		{
			filepath.Join(tarDir, "both.tar"),
			[][]string{
				{
					"Return-Path: <out-582911-B2C71BD37AF148CE9D728B61264F854D@mail.beliefnet.com>",
					"X-Original-To: beliefnet@cp.monitor1.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 447A219825C",
					"\tfor <beliefnet@cp.monitor1.returnpath.net>; Fri,  1 Apr 2011 10:32:42 -0600 (MDT)",
				},
				{
					"From: \"Darty\" <infos@contact-darty.com>",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
				},
			},
		},
		{
			filepath.Join(tarDir, "dirdepth2.tar"),
			[][]string{
				{
					"Return-Path: <out-582911-B2C71BD37AF148CE9D728B61264F854D@mail.beliefnet.com>",
					"X-Original-To: beliefnet@cp.monitor1.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 447A219825C",
					"\tfor <beliefnet@cp.monitor1.returnpath.net>; Fri,  1 Apr 2011 10:32:42 -0600 (MDT)",
				},
				{
					"From: \"Darty\" <infos@contact-darty.com>",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
				},
			},
		},
	}

	for _, c := range cases {
		if _, out := Tar(c.tarPath); !reflect.DeepEqual(out, c.parsedHeaders) {
			t.Errorf("%v returned %v, wanted %v", c.tarPath, out, c.parsedHeaders)
		}
	}
}

func TestCreateArchiveName(t *testing.T) {
	cases := []struct {
		in string
		want string
	}{
		{"test.tar.gz", "test.tar"},
		{"t.tar.gz", "t.tar"},
		{"test.tar.gz.gz", "test.tar.gz"},
	}

	for _, c := range cases {
		if out := CreateArchiveName(c.in); out != c.want {
			t.Errorf("%v returned %v, wanted %v", c.in, out, c.want)
		}
	}
}

