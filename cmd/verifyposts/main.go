package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labudzinski/labudzinski.com/internal/blogpost"
)

func main() {
	repo := flag.String("repo", "", "Ścieżka do repozytorium Hugo (domyślnie: bieżący katalog lub BLOG_REPO)")
	flag.Parse()

	root, err := blogpost.ResolveRepoRoot(*repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "verifyposts: %v\n", err)
		os.Exit(1)
	}

	result, err := blogpost.VerifyPosts(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "verifyposts: %v\n", err)
		for _, verifyErr := range result.Errors {
			fmt.Fprintf(os.Stderr, "  - %v\n", verifyErr)
		}
		os.Exit(1)
	}

	fmt.Printf("verified: %d\n", result.Signed)
}
