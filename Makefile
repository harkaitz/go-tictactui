DESTDIR=
PREFIX=/usr/local
all:
clean:
install:
## -- AUTO-GO --
GO_PROGRAMS += bin/tictactoe$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
deps:
bin/tictactoe$(EXE): deps 
	go build -o $@ $(TICTACTOE_FLAGS) $(GO_CONF) ./cmd/tictactoe
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/tictactoe$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
## -- license --
install: install-license
install-license: LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
## -- license --
## -- AUTO-SERVICE --

## -- AUTO-SERVICE --
