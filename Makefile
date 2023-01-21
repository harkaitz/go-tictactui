DESTDIR=
PREFIX=/usr/local
all:
clean:
install:
## -- license --
install: install-license
install-license: LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/go-tictactui
## -- license --
## -- AUTO-GO --
GO_PROGRAMS = bin/tictactoe$(EXE) 
.PHONY all-go: $(GO_PROGRAMS)
all:     all-go
install: install-go
clean:   clean-go
bin/tictactoe$(EXE): 
	go build -o $@ $(TICTACTOE_FLAGS) $(GO_CONF) ./cmd/tictactoe
install-go: all-go
	install -d $(DESTDIR)$(PREFIX)/bin
	cp bin/tictactoe$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(GO_PROGRAMS)
## -- AUTO-GO --
