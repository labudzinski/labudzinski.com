package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labudzinski/labudzinski.com/internal/blogpost"
)

func main() {
	repo := flag.String("repo", "", "Ścieżka do repozytorium Hugo (domyślnie: bieżący katalog lub BLOG_REPO)")
	email := flag.String("email", "dominik@labudzinski.com", "Adres e-mail klucza PGP")
	dryRun := flag.Bool("dry-run", false, "Podpisz bez zapisu plików")
	skipDrafts := flag.Bool("skip-drafts", false, "Pomiń posty ze statusem draft")
	flag.Parse()

	root, err := blogpost.ResolveRepoRoot(*repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "signposts: %v\n", err)
		os.Exit(1)
	}

	if err := blogpost.ImportPGPKeyFromEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "signposts: %v\n", err)
		os.Exit(1)
	}

	result, err := blogpost.SignPosts(blogpost.SignOptions{
		RepoRoot:   root,
		Email:      *email,
		Passphrase: os.Getenv("PGP_PASSPHRASE"),
		DryRun:     *dryRun,
		SkipDrafts: *skipDrafts,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "signposts: %v\n", err)
		for _, signErr := range result.Errors {
			fmt.Fprintf(os.Stderr, "  - %v\n", signErr)
		}
		os.Exit(1)
	}

	fmt.Printf("signed: %d, unchanged: %d\n", result.Signed, result.Skipped)
}
