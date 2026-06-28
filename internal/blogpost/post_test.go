package blogpost

import "testing"

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
