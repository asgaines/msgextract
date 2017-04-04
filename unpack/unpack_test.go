package unpack

import (
	"testing"
	"os"
	"reflect"
	"io/ioutil"
	"path/filepath"
)

func TestUnGzipCreatesArchive(t *testing.T) {
	fileName := "testEmails.tar.gz"
	unGzipDir := "../test_files/targzs"
	tmpDir, err := ioutil.TempDir("../test_files", "tmp")
	if err != nil {
		t.Error(err)
	}

	// Target path where contents of file to be stored 
	// after unpacking
	unzippedPath := CreateArchiveName(tmpDir, fileName)
	// Gzipped filename will be same as unzipped archive with
	// .gz extension added
	zippedPath := filepath.Join(unGzipDir, fileName)

	// Ensure unpacked tar file does not yet exist
	if _, err := os.Open(unzippedPath); !os.IsNotExist(err) {
		// The file exists
		t.Errorf("Unzipped file already exists: %v", unzippedPath)
	}

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
		headerLines [][]string
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
		{
			filepath.Join(tarDir, "full_file.tar"),
			[][]string{
				{
					"Return-Path: <infos@contact-darty.com>",
					"X-Original-To: 1000mercis@cp.assurance.returnpath.net",
					"Delivered-To: assurance@localhost.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 426E5198271",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri,  1 Apr 2011 08:17:45 -0600 (MDT)",
					"Received: from smtp-8-ft1.mm.fr.colt.net (smtp-7-ft1.mm.fr.colt.net [62.23.8.162])",
					"\tby mxa-d1.returnpath.net (Postfix) with ESMTP id 2906A1CD",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri,  1 Apr 2011 08:17:44 -0600 (MDT)",
					"Received: from host.25.62.23.62.rev.coltfrance.com ([62.23.62.25]:62162 helo=contact-darty.com)",
					"\tby massmail-ft1.infra.coltfrance.com with esmtp (Exim)",
					"\tid 1Q5fAU-00030S-4i",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri, 01 Apr 2011 16:17:42 +0200",
					"From: \"Darty\" <infos@contact-darty.com>",
					"To: 1000mercis@cp.assurance.returnpath.net",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
					"Message-ID: <20110401161739.E3786358A9D7B977@contact-darty.com>",
					"MIME-Version: 1.0",
					"x-idmail: DartyCRM_322_385774_10000",
					"Content-Type: text/html;",
					"\tcharset=\"iso-8859-1\"",
					"Content-Transfer-Encoding: 7bit",
				},
			},
		},
		{
			filepath.Join(tarDir, "big.tar"),
			[][]string{
				{
					"Return-Path: <infos@contact-darty.com>",
					"X-Original-To: 1000mercis@cp.assurance.returnpath.net",
					"Delivered-To: assurance@localhost.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 426E5198271",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri,  1 Apr 2011 08:17:45 -0600 (MDT)",
					"Received: from smtp-8-ft1.mm.fr.colt.net (smtp-7-ft1.mm.fr.colt.net [62.23.8.162])",
					"\tby mxa-d1.returnpath.net (Postfix) with ESMTP id 2906A1CD",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri,  1 Apr 2011 08:17:44 -0600 (MDT)",
					"Received: from host.25.62.23.62.rev.coltfrance.com ([62.23.62.25]:62162 helo=contact-darty.com)",
					"\tby massmail-ft1.infra.coltfrance.com with esmtp (Exim)",
					"\tid 1Q5fAU-00030S-4i",
					"\tfor <1000mercis@cp.assurance.returnpath.net>; Fri, 01 Apr 2011 16:17:42 +0200",
					"From: \"Darty\" <infos@contact-darty.com>",
					"To: 1000mercis@cp.assurance.returnpath.net",
					"Subject: Cuit Vapeur 29.90 euros, Nintendo 3DS 239 euros, GPS TOM TOM 139 euros... decouvrez VITE tous les bons plans du weekend !",
					"Date: 01 Apr 2011 16:17:41 +0200",
					"Message-ID: <20110401161739.E3786358A9D7B977@contact-darty.com>",
					"MIME-Version: 1.0",
					"x-idmail: DartyCRM_322_385774_10000",
					"Content-Type: text/html;",
					"\tcharset=\"iso-8859-1\"",
					"Content-Transfer-Encoding: 7bit",
				},
				{
					"Return-Path: <survey@mindspaymails.com>",
					"X-Original-To: aamarketinginc@cp.monitor1.returnpath.net",
					"Delivered-To: assurance@localhost.returnpath.net",
					"Received: from mxa-d1.returnpath.net (unknown [10.8.2.117])",
					"\tby cpa-d1.returnpath.net (Postfix) with ESMTP id 2DC9B19825C",
					"\tfor <aamarketinginc@cp.monitor1.returnpath.net>; Thu, 31 Mar 2011 22:19:53 -0600 (MDT)",
					"Received: from mail01.mindspaymails.com (mail01.mindspaymails.com [72.3.153.202])",
					"\tby mxa-d1.returnpath.net (Postfix) with ESMTP id E6B1CD85",
					"\tfor <aamarketinginc@cp.monitor1.returnpath.net>; Thu, 31 Mar 2011 22:19:52 -0600 (MDT)",
					"Received: from mail01.incentiverabbit.com (127.0.0.1) by mail01.mindspaymails.com id hil96g0mafs6 for <aamarketinginc@cp.monitor1.returnpath.net>; Thu, 31 Mar 2011 23:19:52 -0500 (envelope-from <survey@mindspaymails.com>)",
					"From: MindsPay<survey@mindspaymails.com> ",
					"To: aamarketinginc@cp.monitor1.returnpath.net",
					"Date: Thu, 31 Mar 2011 23:19:52 -0500",
					"Subject:Paid Mail : Offer #10491 get $4.00",
					"MIME-Version: 1.0",
					"Message-ID:<MP1301631592801EH10491@mindspay.com>",
					"X-BBounce:MP1301631592801EH10491&userid=93321&email=aamarketinginc@cp.monitor1.returnpath.net&",
					"X-campaignid:MP20110331.10491",
					"List-Unsubscribe:<mailto:unsubscribe@mindspay.com?body=userid_93321-email_aamarketinginc@cp.monitor1.returnpath.net&subject=Unsubscribe_From_MindsPay_Instantly>",
					"Content-Type: text/html; charset=iso-8859-1",
					"Content-Transfer-Encoding: 8bit",
				},
			},
		},
	}

	for _, c := range cases {
		headerChan := make(chan []string, len(c.headerLines))

		go func() {
			err := Tar(c.tarPath, headerChan)
			if err != nil {
				t.Error(err)
			}
			close(headerChan)
		}()

		var receivedHeaders [][]string
		numHeaders := 0
		for headers := range headerChan {
			receivedHeaders = append(receivedHeaders, headers)
			numHeaders++
		}

		if !reflect.DeepEqual(receivedHeaders, c.headerLines) {
			t.Errorf("Channel received %v, wanted %v", receivedHeaders, c.headerLines)
		}

		if numHeaders != len(c.headerLines) {
			t.Errorf("Channel received %v headers, wanted %v", numHeaders, len(c.headerLines))
		}
	}
}

func TestCreateArchiveName(t *testing.T) {
	tmpDir, err := ioutil.TempDir("../test_files", "tmp")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		in string
		want string
	}{
		{"test.tar.gz", filepath.Join(tmpDir, "test.tar")},
		{"a/long/path/to/test.tar.gz", filepath.Join(tmpDir, "test.tar")},
		{"t.tar.gz", filepath.Join(tmpDir, "t.tar")},
		{"test.tar.gz.gz", filepath.Join(tmpDir, "test.tar.gz")},
	}

	for _, c := range cases {
		if out := CreateArchiveName(tmpDir, c.in); out != c.want {
			t.Errorf("%v returned %v, wanted %v", c.in, out, c.want)
		}
	}
}

