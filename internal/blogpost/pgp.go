package blogpost

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const defaultPGPEmail = "dominik@labudzinski.com"

type SignOptions struct {
	RepoRoot    string
	Email       string
	Passphrase  string
	DryRun      bool
	SkipDrafts  bool
}

type SignResult struct {
	Signed  int
	Skipped int
	Errors  []error
}

func SignPosts(opts SignOptions) (SignResult, error) {
	if strings.TrimSpace(opts.RepoRoot) == "" {
		return SignResult{}, fmt.Errorf("repo root is required")
	}
	if strings.TrimSpace(opts.Email) == "" {
		opts.Email = defaultPGPEmail
	}

	postsDir := filepath.Join(opts.RepoRoot, "content", "posts")
	entries, err := os.ReadDir(postsDir)
	if err != nil {
		return SignResult{}, fmt.Errorf("read posts dir: %w", err)
	}

	fingerprint, err := gpgFingerprint(opts.Email)
	if err != nil {
		return SignResult{}, fmt.Errorf("pgp fingerprint: %w", err)
	}

	var result SignResult
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		path := filepath.Join(postsDir, entry.Name())
		changed, err := signPostFile(path, opts, fingerprint)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Errorf("%s: %w", entry.Name(), err))
			continue
		}
		if changed {
			result.Signed++
		} else {
			result.Skipped++
		}
	}

	if len(result.Errors) > 0 {
		return result, fmt.Errorf("%d post(s) failed to sign", len(result.Errors))
	}
	return result, nil
}

func VerifyPosts(repoRoot string) (SignResult, error) {
	if strings.TrimSpace(repoRoot) == "" {
		return SignResult{}, fmt.Errorf("repo root is required")
	}

	postsDir := filepath.Join(repoRoot, "content", "posts")
	entries, err := os.ReadDir(postsDir)
	if err != nil {
		return SignResult{}, fmt.Errorf("read posts dir: %w", err)
	}

	var result SignResult
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		path := filepath.Join(postsDir, entry.Name())
		ok, err := verifyPostFile(path)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Errorf("%s: %w", entry.Name(), err))
			continue
		}
		if ok {
			result.Signed++
		} else {
			result.Skipped++
		}
	}

	if len(result.Errors) > 0 {
		return result, fmt.Errorf("%d post(s) failed verification", len(result.Errors))
	}
	return result, nil
}

func verifyPostFile(path string) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	fm, body, err := splitFrontMatter(raw, false)
	if err != nil {
		return false, err
	}

	fingerprint, _ := fm["pgp_key_fingerprint"].(string)
	signature, _ := fm["pgp_signature"].(string)
	if strings.TrimSpace(fingerprint) == "" || strings.TrimSpace(signature) == "" {
		return false, fmt.Errorf("missing pgp_key_fingerprint or pgp_signature")
	}

	ascPath := path + ".asc"
	ascData, err := os.ReadFile(ascPath)
	if err != nil {
		return false, fmt.Errorf("read signature file: %w", err)
	}
	if strings.TrimSpace(string(ascData)) != strings.TrimSpace(signature) {
		return false, fmt.Errorf("front matter signature differs from %s", filepath.Base(ascPath))
	}

	canonical := canonicalBody(body)
	if !gpgVerify(canonical, ascData, os.Getenv("PGP_PASSPHRASE")) {
		return false, fmt.Errorf("gpg verify failed for signed payload")
	}
	return true, nil
}

func signPostFile(path string, opts SignOptions, fingerprint string) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}

	fm, body, err := splitFrontMatter(raw, true)
	if err != nil {
		return false, err
	}

	if opts.SkipDrafts && isDraft(fm) {
		return false, nil
	}

	canonical := canonicalBody(body)
	ascPath := path + ".asc"

	if signatureMatches(fm, fingerprint, canonical, ascPath, opts) {
		return false, nil
	}

	signature, err := gpgDetachSign(canonical, opts.Email, opts.Passphrase)
	if err != nil {
		return false, err
	}

	fm["pgp_key_fingerprint"] = fingerprint
	fm["pgp_signature"] = strings.TrimSpace(signature) + "\n"

	updated, err := joinFrontMatter(fm, body)
	if err != nil {
		return false, err
	}

	if opts.DryRun {
		return true, nil
	}

	if err := os.WriteFile(path, updated, 0o644); err != nil {
		return false, err
	}
	if err := os.WriteFile(ascPath, []byte(signature), 0o644); err != nil {
		return false, err
	}
	return true, nil
}

func isDraft(fm map[string]interface{}) bool {
	switch v := fm["draft"].(type) {
	case bool:
		return v
	case string:
		return strings.EqualFold(strings.TrimSpace(v), "true")
	default:
		return false
	}
}

func signatureMatches(fm map[string]interface{}, fingerprint, body, ascPath string, opts SignOptions) bool {
	existingFP, _ := fm["pgp_key_fingerprint"].(string)
	existingSig, _ := fm["pgp_signature"].(string)
	if strings.TrimSpace(existingFP) == "" || strings.TrimSpace(existingSig) == "" {
		return false
	}
	if strings.TrimSpace(existingFP) != strings.TrimSpace(fingerprint) {
		return false
	}

	ascData, err := os.ReadFile(ascPath)
	if err != nil {
		return false
	}
	if strings.TrimSpace(string(ascData)) != strings.TrimSpace(existingSig) {
		return false
	}

	return gpgVerify(body, ascData, opts.Passphrase)
}

func splitFrontMatter(raw []byte, stripPGP bool) (map[string]interface{}, string, error) {
	text := string(raw)
	if !strings.HasPrefix(text, "---\n") {
		return nil, "", fmt.Errorf("missing front matter")
	}

	rest := text[4:]
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return nil, "", fmt.Errorf("unclosed front matter")
	}

	fmText := rest[:end]
	body := rest[end+4:]
	body = strings.TrimLeft(body, "\n")

	fm := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(fmText), &fm); err != nil {
		return nil, "", fmt.Errorf("parse front matter: %w", err)
	}

	if stripPGP {
		delete(fm, "pgp_key_fingerprint")
		delete(fm, "pgp_signature")
	}

	return fm, body, nil
}

func joinFrontMatter(fm map[string]interface{}, body string) ([]byte, error) {
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(fm); err != nil {
		return nil, fmt.Errorf("encode front matter: %w", err)
	}
	if err := enc.Close(); err != nil {
		return nil, err
	}

	var out strings.Builder
	out.WriteString("---\n")
	out.WriteString(strings.TrimSpace(buf.String()))
	out.WriteString("\n---\n\n")
	out.WriteString(strings.TrimRight(body, "\n"))
	out.WriteString("\n")
	return []byte(out.String()), nil
}

func canonicalBody(body string) string {
	return strings.TrimRight(body, "\n") + "\n"
}

func gpgFingerprint(email string) (string, error) {
	out, err := exec.Command("gpg", "--batch", "--with-colons", "--fingerprint", email).Output()
	if err != nil {
		return "", fmt.Errorf("gpg fingerprint for %q: %w", email, err)
	}

	for _, line := range strings.Split(string(out), "\n") {
		if !strings.HasPrefix(line, "fpr:") {
			continue
		}
		parts := strings.Split(line, ":")
		if len(parts) < 10 || parts[9] == "" {
			continue
		}
		return formatFingerprint(parts[9]), nil
	}
	return "", fmt.Errorf("no fingerprint found for %q", email)
}

func formatFingerprint(hex string) string {
	hex = strings.ToUpper(strings.ReplaceAll(hex, " ", ""))
	var parts []string
	for i := 0; i < len(hex); i += 4 {
		end := i + 4
		if end > len(hex) {
			end = len(hex)
		}
		parts = append(parts, hex[i:end])
	}
	return strings.Join(parts, " ")
}

func gpgDetachSign(content, email, passphrase string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "blogpost-sign-*")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmpDir)

	contentPath := filepath.Join(tmpDir, "content.md")
	if err := os.WriteFile(contentPath, []byte(content), 0o600); err != nil {
		return "", err
	}

	args := []string{
		"--batch", "--yes", "--armor", "--detach-sign",
		"--local-user", email, "--output", "-", contentPath,
	}
	if passphrase != "" {
		args = append([]string{"--pinentry-mode", "loopback", "--passphrase-fd", "0"}, args...)
	}

	cmd := exec.Command("gpg", args...)
	if passphrase != "" {
		cmd.Stdin = strings.NewReader(passphrase + "\n")
	}

	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("gpg detach-sign: %w", err)
	}
	return string(out), nil
}

func gpgVerify(content string, asc []byte, passphrase string) bool {
	tmpDir, err := os.MkdirTemp("", "blogpost-pgp-*")
	if err != nil {
		return false
	}
	defer os.RemoveAll(tmpDir)

	contentPath := filepath.Join(tmpDir, "content.md")
	sigPath := filepath.Join(tmpDir, "content.md.asc")
	if err := os.WriteFile(contentPath, []byte(content), 0o600); err != nil {
		return false
	}
	if err := os.WriteFile(sigPath, asc, 0o600); err != nil {
		return false
	}

	cmd := exec.Command("gpg", "--batch", "--verify", sigPath, contentPath)
	if passphrase != "" {
		cmd = exec.Command("gpg", "--batch", "--pinentry-mode", "loopback",
			"--passphrase-fd", "0", "--verify", sigPath, contentPath)
		cmd.Stdin = strings.NewReader(passphrase + "\n")
	}

	return cmd.Run() == nil
}

func ImportPGPKeyFromEnv() error {
	key := strings.TrimSpace(os.Getenv("PGP_PRIVATE_KEY"))
	if key == "" {
		return nil
	}

	cmd := exec.Command("gpg", "--batch", "--import")
	cmd.Stdin = strings.NewReader(key)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("import pgp key: %w", err)
	}
	return nil
}
