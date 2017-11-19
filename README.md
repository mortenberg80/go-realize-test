# go-realize-test

Project created to figure out why realize does not work with cmd / pkg-structure for projects.


### Setup
This project uses [dep](https://github.com/golang/dep#setup) for dependency management.

Install project dependencies with either
```
# Make target
make deps

# or dep command
dep ensure
```

Realize
```
go get github.com/tockins/realize
```

Build project with: `make`

Realize _should_ work with `make dev`, but this seems to fail at the moment.

### Error report

#### OS X

OS: Mac OS X El Capitan Version 10.11.6

```
$ uname -a
Darwin macberg2 15.6.0 Darwin Kernel Version 15.6.0: Sun Jun  4 21:43:07 PDT 2017; root:xnu-3248.70.3~1/RELEASE_X86_64 x86_64
```

GO-version: `go version go1.9.2 darwin/amd64`

Realize version: `1.5.2`

Ouput:
```
$ make dev
realize start
[09:29:25][GO-REALIZE-TEST] : Watching 4 file/s 7 folder/s
[09:29:25][GO-REALIZE-TEST] : Building...
[09:29:26][GO-REALIZE-TEST] : Built in 0.912 s
[09:29:26][GO-REALIZE-TEST] : Running..
[09:29:26][GO-REALIZE-TEST] : Build not found
```

Generated realize.yaml-file:

```yaml
settings:
  files:
    outputs:
      status: true
      name: outputs.log
    logs:
      status: true
      name: logs.log
    errors:
      status: true
      name: errors.log
  legacy:
    force: false
    interval: 100ms
server:
  status: false
  open: false
  port: 5001
  host: localhost
schema:
- name: go-realize-test
  path: /Users/mortenberg/go/src/github.com/mortenberg80/go-realize-test
  commands:
    build:
      method: go build -o bin/go-realize-test ./cmd/go-realize-test
      status: true
    run: true
  watcher:
    paths:
    - /
    extensions:
    - go
    ignored_paths:
    - .git
    - .realize
    - db
    - bin
    - etc
    - support
    - vendor
```

#### Ubuntu 16.04

I get the same problem on my Ubuntu machine, with the same Go and realize versions.
