LNSC=lnsc

help:
	@echo make test-lns

test-lns:
	$(LNSC) test.lns save -langGo --main test
	$(LNSC) test.lns mkmain main.go
	go build
	rm -f hoge.sqlite3
	./test
