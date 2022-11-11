PICKLE = ../go-heapdump.pk
EXAMPLE_HEAPDUMP = example.heapdump

all: build

build: $(EXAMPLE_HEAPDUMP)

build-in-docker:
	docker run --rm -v $(PWD):/s -w /s golang:alpine sh -c 'apk add --no-cache make && make'

clean:
	rm -f $(EXAMPLE_HEAPDUMP) example

test: $(EXAMPLE_HEAPDUMP) $(PICKLE)
	poke -l $(PICKLE) -c '.file $(EXAMPLE_HEAPDUMP)' -s test.pk

example: example.go
	go build -o $@ $<

$(EXAMPLE_HEAPDUMP): example
	./example $@

.PHONY: all build build-in-docker clean test
