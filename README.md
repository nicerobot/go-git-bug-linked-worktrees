# Show bug

    git clone git@github.com:nicerobot/go-git-bug-linked-worktrees
    cd go-git-bug-linked-worktrees
    make

# Output

    $ make
    git clone git@github.com:src-d/go-git.git
    Cloning into 'go-git'...
    cd go-git; git fetch --all
    go-git-bug/go-git
    Fetching origin
    cd go-git; git worktree add -b new-ref ../go-git-linked/
    go-git-bug/go-git
    Preparing worktree (new branch 'new-ref')
    HEAD is now at 1a7db85bca70 Merge pull request #1231 from alexandear/fix-typos
    go build -o bugged


    ---- Works as expected
    go-git-bug/go-git
    git rev-parse --git-dir
    .git
    git rev-parse HEAD
    1a7db85bca7027d90afdb5ce711622aaac9feaed
    2019/12/01 11:39:36 worktree: go-git-bug/go-git
    2019/12/01 11:39:36 1a7db85bca7027d90afdb5ce711622aaac9feaed refs/heads/master
    2019/12/01 11:39:36 1a7db85bca7027d90afdb5ce711622aaac9feaed refs/heads/new-ref
    2019/12/01 11:39:36 head: 1a7db85bca7027d90afdb5ce711622aaac9feaed refs/heads/master


    ---- Bug this does not work as expected
    go-git-bug/go-git-linked
    git rev-parse --git-dir
    go-git-bug/go-git/.git/worktrees/go-git-linked
    git rev-parse HEAD
    1a7db85bca7027d90afdb5ce711622aaac9feaed
    2019/12/01 11:39:36 worktree: go-git-bug/go-git-linked
    2019/12/01 11:39:36 head: reference not found
    make: *** [bug] Error 1

# Call-stack

    repo, err := git.PlainOpenWithOptions(".", &git.PlainOpenOptions{DetectDotGit: true})
    repo.Head()
    storer.ResolveReference(r.Storer, plumbing.HEAD)
    resolveReference(s, r, 0)
    s.Reference(r.Target())
    r.dir.Ref(n) // <- `n` (i.e. `r.Target()`) is `ref/heads/new-ref`
    d.readReferenceFile(".", name.String())
    d.fs.Open(path) // <- this `path` is wrong for this situation

