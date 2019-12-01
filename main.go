package main

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"log"
)

//
func main() {
	repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		log.Fatalf("open: %s", err)
	}
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("%s", err)
	}
	worktreeRoot := worktree.Filesystem.Root()
	if worktreeRoot == "" {
		log.Fatalf("worktree: %s", err)
	}
	log.Printf("worktree: %s", worktreeRoot)
	b, err := repo.Branches()
	if err != nil {
		log.Fatalf("branches: %s", err)
	}
	err = b.ForEach(func(reference *plumbing.Reference) error {
		log.Printf("%+v", reference)
		return nil
	})
	if err != nil {
		log.Fatalf("each branch: %s", err)
	}
	head, err := repo.Head()
	if err != nil {
		log.Fatalf("head: %s", err)
	}
	log.Printf("head: %s", head.String())
}
