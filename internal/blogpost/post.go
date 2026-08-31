package blogpost

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Post struct {
	Title       string
	Content     string
	Description string
	Tags        []string
	Draft       bool
}

const metaDescriptionLimit = 160

var (
	slugSanitizer  = regexp.MustCompile(`[^a-z0-9\-]+`)
	markdownInline = regexp.MustCompile("`+|[*_]+")
)

func Slugify(title string) string {
	title = strings.ToLower(strings.TrimSpace(title))
	title = stripDiacritics(title)
	title = strings.ReplaceAll(title, " ", "-")
	title = slugSanitizer.ReplaceAllString(title, "-")
	title = strings.Trim(title, "-")
	for strings.Contains(title, "--") {
		title = strings.ReplaceAll(title, "--", "-")
	}
	if title == "" {
		title = "post"
	}
	return title
}

func stripDiacritics(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, _ := transform.String(t, s)
	return out
}

func WritePost(repoRoot string, post Post) (string, error) {
	if strings.TrimSpace(post.Title) == "" {
		return "", fmt.Errorf("tytuł nie może być pusty")
	}

	slug := Slugify(post.Title)
	filename := fmt.Sprintf("%s.md", slug)
	dir := filepath.Join(repoRoot, "content", "posts")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); err == nil {
		path = filepath.Join(dir, fmt.Sprintf("%s-%s.md", slug, time.Now().Format("20060102-150405")))
	}

	content := renderFrontMatter(post, slug) + strings.TrimSpace(post.Content) + "\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", err
	}
	return path, nil
}

func DeriveDescription(explicit, content string) string {
	src := strings.TrimSpace(explicit)
	if src == "" {
		src = firstPlainParagraph(content)
	}
	return clampMetaDescription(src)
}

func firstPlainParagraph(md string) string {
	s := strings.ReplaceAll(md, "\r\n", "\n")
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	var parts []string
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			if len(parts) > 0 {
				break
			}
			continue
		}
		if len(parts) == 0 {
			if strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "---") {
				continue
			}
		}
		parts = append(parts, trimmed)
	}
	plain := strings.Join(parts, " ")
	plain = markdownInline.ReplaceAllString(plain, "")
	return strings.TrimSpace(plain)
}

func clampMetaDescription(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	if s == "" {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= metaDescriptionLimit {
		return s
	}
	cut := runes[:metaDescriptionLimit]
	if i := lastSpace(cut); i > metaDescriptionLimit/2 {
		cut = cut[:i]
	}
	return strings.TrimSpace(string(cut)) + "…"
}

func lastSpace(runes []rune) int {
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsSpace(runes[i]) {
			return i
		}
	}
	return -1
}

func renderFrontMatter(post Post, slug string) string {
	desc := DeriveDescription(post.Description, post.Content)
	var b strings.Builder
	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: %q\n", post.Title)
	fmt.Fprintf(&b, "date: %s\n", time.Now().Format("2006-01-02T15:04:05-07:00"))
	fmt.Fprintf(&b, "draft: %t\n", post.Draft)
	if slug != "" {
		fmt.Fprintf(&b, "slug: %s\n", slug)
	}
	if desc != "" {
		fmt.Fprintf(&b, "description: %q\n", desc)
	}
	b.WriteString("toc: false\n")
	b.WriteString("tags:\n")
	for _, tag := range post.Tags {
		tag = strings.TrimSpace(tag)
		if tag != "" {
			fmt.Fprintf(&b, "  - %s\n", tag)
		}
	}
	b.WriteString("---\n\n")
	return b.String()
}

func LoadExistingTags(repoRoot string) ([]string, error) {
	postsDir := filepath.Join(repoRoot, "content", "posts")
	entries, err := os.ReadDir(postsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	seen := make(map[string]struct{})
	var tags []string

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		fileTags, err := parseTagsFromFile(filepath.Join(postsDir, entry.Name()))
		if err != nil {
			continue
		}
		for _, tag := range fileTags {
			key := strings.ToLower(tag)
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			tags = append(tags, tag)
		}
	}
	return tags, nil
}

func parseTagsFromFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() || strings.TrimSpace(scanner.Text()) != "---" {
		return nil, fmt.Errorf("brak front matter")
	}

	var tags []string
	inTags := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "---" {
			break
		}
		if strings.HasPrefix(strings.TrimSpace(line), "tags:") {
			inTags = true
			continue
		}
		if inTags {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "- ") {
				tags = append(tags, strings.TrimSpace(trimmed[2:]))
				continue
			}
			if trimmed != "" && !strings.HasPrefix(trimmed, "#") {
				inTags = false
			}
		}
	}
	return tags, scanner.Err()
}
