package blogpost

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"Backup Schrödingera", "backup-schrodingera"},
		{"Szyfrowanie plików", "szyfrowanie-plikow"},
		{"  Hello World  ", "hello-world"},
	}
	for _, tc := range tests {
		got := Slugify(tc.in)
		if got != tc.want {
			t.Errorf("Slugify(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestFilterTags(t *testing.T) {
	all := []string{"backup", "szyfrowanie", "PHP", "Symfony"}
	selected := map[string]struct{}{"backup": {}}
	got := FilterTags(all, "sz", selected)
	if len(got) != 1 || got[0] != "szyfrowanie" {
		t.Fatalf("FilterTags() = %v", got)
	}
}

func TestApplyTagSuggestion(t *testing.T) {
	got := ApplyTagSuggestion("back, sz", "szyfrowanie")
	want := "back, szyfrowanie, "
	if got != want {
		t.Fatalf("ApplyTagSuggestion() = %q, want %q", got, want)
	}
}

func TestLoadExistingTags(t *testing.T) {
	tags, err := LoadExistingTags("../../")
	if err != nil {
		t.Fatal(err)
	}
	if len(tags) < 3 {
		t.Fatalf("expected at least 3 tags, got %d: %v", len(tags), tags)
	}
}

func TestDeriveDescription(t *testing.T) {
	tests := []struct {
		name, explicit, content, want string
	}{
		{
			name:     "explicit wins",
			explicit: "Krótki opis pod wyszukiwarki.",
			content:  "Długi pierwszy akapit, który nie powinien trafić do meta.",
			want:     "Krótki opis pod wyszukiwarki.",
		},
		{
			name:    "first paragraph",
			content: "Od dawna odkładałem backup na później.\n\nDrugi akapit już nie.",
			want:    "Od dawna odkładałem backup na później.",
		},
		{
			name:    "collapse whitespace",
			content: "Foo   bar\nbaz.",
			want:    "Foo bar baz.",
		},
		{
			name:    "strip heading",
			content: "# Tytuł\n\nWłaściwy lead artykułu o backupie.",
			want:    "Właściwy lead artykułu o backupie.",
		},
		{
			name:    "empty content",
			content: "   ",
			want:    "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DeriveDescription(tc.explicit, tc.content)
			if got != tc.want {
				t.Fatalf("DeriveDescription() = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestDeriveDescriptionTruncatesAt160Runes(t *testing.T) {
	long := strings.Repeat("abecadło ", 40)
	got := DeriveDescription("", long)
	if n := len([]rune(got)); n > 160 {
		t.Fatalf("len=%d, got %q", n, got)
	}
	if !strings.HasSuffix(got, "…") && !strings.HasSuffix(got, "...") {
		t.Fatalf("expected ellipsis, got %q", got)
	}
}

func TestWritePostWritesDescriptionAndSlug(t *testing.T) {
	root := t.TempDir()
	path, err := WritePost(root, Post{
		Title:   "Backup Schrödingera",
		Content: "Wyobraźmy sobie, że każdy robi regularnie kopię zapasową.\n\nDrugi akapit.",
		Tags:    []string{"backup"},
	})
	if err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	if filepath.Base(path) != "backup-schrodingera.md" {
		t.Fatalf("filename = %s", filepath.Base(path))
	}
	if !strings.Contains(text, "slug: backup-schrodingera\n") {
		t.Fatalf("missing slug:\n%s", text)
	}
	if !strings.Contains(text, `description: "Wyobraźmy sobie, że każdy robi regularnie kopię zapasową."`) {
		t.Fatalf("missing description:\n%s", text)
	}
}
