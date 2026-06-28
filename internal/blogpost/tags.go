package blogpost

import (
	"sort"
	"strings"
)

func FilterTags(all []string, query string, selected map[string]struct{}) []string {
	query = strings.ToLower(strings.TrimSpace(query))
	parts := splitTagInput(query)
	last := ""
	if len(parts) > 0 {
		last = parts[len(parts)-1]
	}

	var matches []string
	for _, tag := range all {
		if _, used := selected[strings.ToLower(tag)]; used {
			continue
		}
		if last == "" || strings.HasPrefix(strings.ToLower(tag), last) {
			matches = append(matches, tag)
		}
	}
	sort.Strings(matches)
	return matches
}

func ParseSelectedTags(input string) []string {
	parts := splitTagInput(input)
	var tags []string
	seen := make(map[string]struct{})
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		key := strings.ToLower(part)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		tags = append(tags, part)
	}
	return tags
}

func CurrentTagFragment(input string) string {
	parts := splitTagInput(input)
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func ApplyTagSuggestion(input, suggestion string) string {
	parts := splitTagInput(input)
	if len(parts) == 0 {
		return suggestion + ", "
	}
	parts[len(parts)-1] = suggestion
	return strings.Join(parts, ", ") + ", "
}

func splitTagInput(input string) []string {
	raw := strings.Split(input, ",")
	var parts []string
	for _, part := range raw {
		part = strings.TrimSpace(part)
		if part != "" {
			parts = append(parts, part)
		}
	}
	return parts
}

func SelectedTagSet(tags []string) map[string]struct{} {
	set := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		set[strings.ToLower(tag)] = struct{}{}
	}
	return set
}
