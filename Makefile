DESTDIR=
PREFIX=/usr/local
all:
clean:
install:
## -- license --
ifneq ($(PREFIX),)
install: install-license
install-license: LICENSE
	@echo 'I share/doc/go-tictactui/LICENSE'
	@mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
	@cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
endif
## -- license --
## -- AUTO-GO --
all:     all-go
install: install-go
clean:   clean-go
all-go:
	@echo "B bin/tictactoe   ./cmd/tictactoe"
	@go build -o bin/tictactoe   ./cmd/tictactoe
install-go: all-go
	@install -d $(DESTDIR)$(PREFIX)/bin
	@echo I bin/tictactoe
	@cp bin/tictactoe   $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f bin/tictactoe
## -- AUTO-GO --
