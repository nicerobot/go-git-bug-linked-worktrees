.PHONY: bug help

THIS := $(abspath $(lastword $(MAKEFILE_LIST)))
THISD := $(patsubst %/,%,$(dir $(THIS)))
HELP := $(MAKE) -f $(THIS) help
BUGGED := $(THISD)/bugged

bug: go-git go-git-linked bugged
	@echo '\n\n---- Works as expected'; cd go-git; $(HELP); $(BUGGED)
	@echo '\n\n---- Bug this does not work as expected'; cd go-git-linked; $(HELP); $(BUGGED)

bugged:
	go build -o bugged

go-git: go-git/go.mod
	cd $@; git fetch --all

go-git/go.mod:
	git clone git@github.com:src-d/go-git.git

go-git-linked: go-git-linked/go.mod

go-git-linked/go.mod:
	cd go-git; git worktree add -b new-ref ../$(dir $@)

help:
	git rev-parse --git-dir
	git rev-parse HEAD

clean:
	@rm -rf bugged || true
	@rm -rf go-git || true
	@rm -rf go-git-linked || true
