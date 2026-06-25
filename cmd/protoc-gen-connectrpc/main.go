package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-connect-py/internal/runner"
)

func main() {
	os.Exit(runner.Run("protoc-gen-connectrpc", os.Args[1:], os.Stdin, os.Stdout, os.Stderr, "."))
}
