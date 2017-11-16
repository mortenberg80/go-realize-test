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

