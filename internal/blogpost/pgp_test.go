package blogpost

import (
	"strings"
	"testing"
)

func TestSplitFrontMatterStripsPGPWhenSigning(t *testing.T) {
	raw := []byte(`---
title: "Backup strategy"
pgp_key_fingerprint: "OLD"
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----
  abc
  -----END PGP SIGNATURE-----
tags:
  - backup
---

Body line one.
`)

	fm, body, err := splitFrontMatter(raw, true)
	if err != nil {
		t.Fatal(err)
	}
	if fm["title"] != "Backup strategy" {
		t.Fatalf("title = %v", fm["title"])
	}
	if _, ok := fm["pgp_key_fingerprint"]; ok {
		t.Fatal("pgp_key_fingerprint should be stripped")
	}
	if body != "Body line one.\n" {
		t.Fatalf("body = %q", body)
	}
}

func TestIsPostSigned(t *testing.T) {
	fm := map[string]interface{}{
		"pgp_key_fingerprint": "A1B2",
		"pgp_signature":       "sig",
	}
	if isPostSigned("missing.md", fm) {
		t.Fatal("expected unsigned without .asc file")
	}
}

func TestJoinFrontMatterIncludesSignature(t *testing.T) {
	fm := map[string]interface{}{
		"title":               "Backup strategy",
		"date":                "2026-06-30",
		"pgp_key_fingerprint": "A1B2 C3D4",
		"pgp_signature":       "-----BEGIN PGP SIGNATURE-----\nabc\n-----END PGP SIGNATURE-----\n",
	}

	out, err := joinFrontMatter(fm, "Post body.\n")
	if err != nil {
		t.Fatal(err)
	}
	text := string(out)
	if !strings.Contains(text, "pgp_key_fingerprint:") ||
		!strings.Contains(text, "pgp_signature:") ||
		!strings.Contains(text, "Post body.") ||
		!strings.Contains(text, "---") {
		t.Fatalf("unexpected output:\n%s", text)
	}
}

func TestCanonicalBody(t *testing.T) {
	got := canonicalBody("hello\n\nworld\n\n")
	want := "hello\n\nworld\n"
	if got != want {
		t.Fatalf("canonicalBody() = %q, want %q", got, want)
	}
}

func TestFormatFingerprint(t *testing.T) {
	got := formatFingerprint("a1b2c3d4e5f67890abcdef1234567890abcdef12")
	want := "A1B2 C3D4 E5F6 7890 ABCD EF12 3456 7890 ABCD EF12"
	if got != want {
		t.Fatalf("formatFingerprint() = %q, want %q", got, want)
	}
}
