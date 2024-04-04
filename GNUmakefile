.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=go-tictactui
VERSION=1.0.0
PREFIX=/usr/local

install: install-sh
install-sh:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp scr/tictactui $(DESTDIR)$(PREFIX)/bin
## -- BLOCK:go --
build/tictactoe$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/tictactoe
all: build/tictactoe$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/tictactoe$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/tictactoe$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
