package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-connectrpc",
		LibraryRepo: "connectrpc/connect-py",
		GoReleaser:  true,
	})
	boot.Main()
}
