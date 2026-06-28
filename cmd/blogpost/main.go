package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labudzinski/labudzinski.com/internal/blogpost"
)

func main() {
	repo := flag.String("repo", "", "Ścieżka do repozytorium Hugo (domyślnie: bieżący katalog lub BLOG_REPO)")
	draft := flag.Bool("draft", false, "Zapisz post jako szkic (draft: true)")
	flag.Parse()

	root, err := blogpost.ResolveRepoRoot(*repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "blogpost: %v\n", err)
		os.Exit(1)
	}

	app, err := blogpost.NewApp(root, *draft)
	if err != nil {
		fmt.Fprintf(os.Stderr, "blogpost: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "blogpost: %v\n", err)
		os.Exit(1)
	}
}
