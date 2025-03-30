.PHONY: all $(COMMANDS) clean test

GO_LDFLAGS := -s -w

COMMANDS := product-info product-screenshot server
BINARIES := $(addprefix build/, $(COMMANDS))

all: $(BINARIES)

build/%: cmd/%/*.go
	go build -ldflags="$(GO_LDFLAGS)" -o $@ ./cmd/$*
	chmod +x $@

test:
	go test -v ./...

clean:
	rm -rf build/
