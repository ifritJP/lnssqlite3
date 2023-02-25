LNSC=lnsc

ifneq "$(wildcard Makefile.local)" ""
include Makefile.local
endif

help:
	@echo make all

all:
	$(LNSC) lns/sqlite3/base.lns save -langGo --package sqlite3
	$(LNSC) lns/sqlite3/access.lns save
