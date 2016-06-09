.PHONY: clean

TARGETS=$(shell find sslide lib -type f -iname "*.go")
BIN=simpleslide
SEMVER=$(shell git describe --always --long)

all: $(BIN)

$(BIN): $(TARGETS)
	go build -o $(BIN) -ldflags "-X github.com/klingtnet/simpleslide/lib.Version=$(SEMVER)" sslide/main.go

run: $(BIN)
	./$(BIN)

clean:
	rm -f $(BIN)
