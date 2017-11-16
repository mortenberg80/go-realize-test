binary=bin/go-realize-test
logfile=log/go-realize-test.log
port=9000

default: $(binary)

$(binary): cmd/go-realize-test/*.go
	mkdir -p bin
	go build -o $@ $^

# Create Realize config from template
.realize/realize.yaml: etc/realize.yaml
	mkdir -p $(dir $@) && sed -e "s#%PROJECT_ROOT%#$(shell pwd)#g" $< > $@

# Requires Realize >= 1.5
# go get -u github.com/tockins/realize
dev: .realize/realize.yaml
	realize start

deps:
	dep ensure

godeps: deps

test:
	go test ./cmd/... ./pkg/...

clean:
	rm -f $(binary)
