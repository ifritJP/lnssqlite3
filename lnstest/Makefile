LNSC=lnsc

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif


help:
	@echo make test-go
	@echo make test-lua

test-go:
	$(LNSC) test.lns save -langGo --main test
	$(LNSC) test.lns mkmain main.go
	go build
	rm -f hoge.sqlite3
	./test

test-lua:
	rm -f hoge.sqlite3
	$(LNSC) test.lns -shebang
