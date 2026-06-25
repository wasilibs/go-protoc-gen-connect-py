# go-protoc-gen-connect-py

go-protoc-gen-connect-py is a distribution of [protoc-gen-connect-py][1], that can be built with Go. It does not actually reimplement any
functionality of protoc-gen-connect-py in Go, instead packaging it with the WASI build of [Python][3], and
executing with the pure Go Wasm runtime [wazero][2]. This means that `go install` or `go run`
can be used to execute it, with no need to rely on separate package managers such as pnpm,
on any platform that Go supports.

See [go-protoc-gen-connect-py](https://github.com/wasilibs/go-protoc-gen-py) for the corresponding module for
`protobuf-py` message stubs.

## Installation

Precompiled binaries are available in the [releases](https://github.com/wasilibs/go-protoc-gen-connect-py/releases).
Alternatively, install the plugin you want using `go install`.

```bash
$ go install github.com/wasilibs/go-protoc-gen-connect-py/cmd/protoc-gen-connectrpc@latest
```

To avoid installation entirely, it can be convenient to use `go run`

```bash
$ go run github.com/wasilibs/go-protoc-gen-connect-py/cmd/protoc-gen-connectrpc@latest .
```

[1]: https://github.com/connectrpc/connect-py
[2]: https://wazero.io/
[3]: https://github.com/brettcannon/cpython-wasi-build
